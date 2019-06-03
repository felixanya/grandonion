package ping

import (
	"context"
	"fmt"
	"os/exec"
	"sync"
	"time"
)

const DefaultPingTimeout = 20 * time.Second

func Ping(ipList ...string) map[string]error {
	return PingWithTimeout(DefaultPingTimeout, ipList...)
}

func PingWithTimeout(timeout time.Duration, ipList ...string) map[string]error {
	if len(ipList) == 0 {
		return nil
	}

	ret := make(map[string]error)
	ch := make(chan map[string]error, len(ipList))
	wg := &sync.WaitGroup{}

	for _, ip := range ipList {
		wg.Add(1)
		go func(ip string) {
			mp := make(map[string]error)
			mp[ip] = ping(ip, timeout)

			ch <- mp
			wg.Done()
		}(ip)
	}
	wg.Wait()

	for i := 0; i < len(ipList); i++ {
		for ip, err := range <- ch {
			ret[ip] = err
		}
	}

	return ret
}

func ping(ip string, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout * time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "ping", ip)
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func fping(ip string, timeout int) error {
	cmd := exec.Command("fping", fmt.Sprintf("-A -t%d %s", timeout, ip))
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
