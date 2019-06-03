package main

import "github.com/looplab/fsm"

func main() {
	fsm.NewFSM("wait_create",
		fsm.Events{
			{Name: "wait_conf", Src: []string{"wait_create", "conf_failed", "creating", "create_failed", "initing", "init_failed", "delivery", "delivery_failed"}, Dst: "wait_create"},
			{Name: "wait_create", Src: []string{"wait_conf", ""}},
		},
		fsm.Callbacks{

		},
     )
}
