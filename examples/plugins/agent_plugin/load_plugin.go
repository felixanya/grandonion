package agent_plugin

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"plugin"
	"time"
)

var defaultPluginPath = []string{"./plugins/", "/var/lib/grandonion/plugins/", "/var/run/grandonion/plugins/"}

type PluginOutPut struct {
	Status  string
	Message string
	Error   string
	Code    int
	Data    interface{}
}

type Plugin interface {
	Load() error
	Available() bool

	Enable() error
	Disable() error

	Execute(ctx context.Context, in ...[]byte) (output []byte, err error)
}

type PluginExecutor struct {
}

func (p *PluginExecutor) Execute(pluginName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	//pluginPath := fmt.Sprintf("./plugins/%s.so", pluginName)
	pluginPath, err := p.findPlugin(pluginName)
	if err != nil {
		return err
	}
	plg, err := plugin.Open(pluginPath)
	if err != nil {
		return err
	}

	sym, err := plg.Lookup("TaskPlugin")
	if err != nil {
		return err
	}
	//var concretePlugin Plugin
	concretePlugin, ok := sym.(Plugin)
	if !ok {
		return errors.New("load plugin error")
	}

	if err := concretePlugin.Load(); err != nil {
		return err
	}

	// enable plugin
	if err := concretePlugin.Enable(); err != nil {
		return err
	}

	go func() {
		input := struct {
			Interval    int     `json:"interval"`
			Content     string  `json:"content"`
		}{
			Interval: 5,
			Content: "bonjour le monte",
		}
		inputBt, err := json.Marshal(&input)
		if err != nil {
			log.Println(err)
			return
		}
		output, err := concretePlugin.Execute(ctx, inputBt)
		if err != nil {
			log.Println(err)
			return
		}
		pluginOutput := &PluginOutPut{}
		if err := json.Unmarshal(output, pluginOutput); err != nil {
			log.Println(err)
			return
		}

		fmt.Printf(">>>output: %+v", pluginOutput)
	}()

	time.Sleep(15 * time.Second)
	log.Printf(">>> available: %t\n", concretePlugin.Available())
	if err := concretePlugin.Disable(); err != nil {
		log.Println(">>>disable plugin error,", err.Error())
		return err
	}

	log.Println(">>>disable plugin success")

	input := struct {
		Interval    int     `json:"interval"`
		Content     string  `json:"content"`
	}{
		Interval: 5,
		Content: "hello world",
	}
	inputBt, err := json.Marshal(&input)
	if err != nil {
		log.Println(err)
		return err
	}
	if _, err := concretePlugin.Execute(ctx, inputBt); err != nil {
		log.Println("Execute error,", err.Error())
	}

	return nil
}

func (p *PluginExecutor) findPlugin(pluginName string) (string, error) {
	for _, dp := range defaultPluginPath {
		_, err := os.Stat(fmt.Sprintf("%s%s.so", dp, pluginName))
		if err == nil || os.IsExist(err) {
			return fmt.Sprintf("%s%s.so", dp, pluginName), nil
		}
	}

	return "", errors.New("plugin not found")
}
