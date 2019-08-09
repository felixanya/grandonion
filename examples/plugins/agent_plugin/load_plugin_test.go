package agent_plugin

import (
    "log"
    "testing"
)

func TestPluginExecutor_Execute(t *testing.T) {
    pe := PluginExecutor{}

    err := pe.Execute("timer_task")
    if err != nil {
        log.Fatal(err)
    }
}
