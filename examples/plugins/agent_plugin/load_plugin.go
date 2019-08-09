package agent_plugin

import "context"

type PluginOutPut struct {
	Status 		string
	Message 	string
	Data 		interface{}
}

type Plugin interface {
	Available() bool

	Enable() error
	Disable() error

	Execute(ctx context.Context, in ...interface{}) (output interface{}, err error)
}

type PluginCollection struct {
	Plugins 	map[string]Plugin
	
}

func (p *PluginCollection) Available(pluginName string) bool {

}
