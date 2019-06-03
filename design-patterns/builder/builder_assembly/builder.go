package builder

import "fmt"

type IAllocator interface {
	ApplyIP()
	GenMac()
	AllocPyHost()
}

type Orchestrator struct {
	allocator IAllocator
}

func NewOrchestrate(a IAllocator) *Orchestrator {
	return &Orchestrator{
		allocator: a,
	}
}

func (orch *Orchestrator) Orchestrate() {
	orch.allocator.AllocPyHost()
	orch.allocator.ApplyIP()
	orch.allocator.GenMac()
}

type Allocator struct {
	netRes NetResource
}

type NetResource struct {
	IP 		string
	PyIP 	string
	IpSeg 	string
	Macs 	[]string
}

func (al *Allocator) ApplyIP() {
	fmt.Println("applying ip address...")
	ip := "192.168.168.11"
	al.netRes.IP = ip
}

func (al *Allocator) GenMac() {
	fmt.Println("generating mac address...")
	al.netRes.Macs = []string{"mad1", "mac2", "mac3"}
}

func (al *Allocator) AllocPyHost() {
	fmt.Println("allocating physical host...")
	al.netRes.PyIP = "192.168.168.10"
	al.netRes.IpSeg = "192.168.168.0/24"
}

func (nr *NetResource) ListAllocated() {
	fmt.Printf("Allocated resource:\n IP:\t%s\n PyIP:\t%s\n Macs:\t%s\n", nr.IP, nr.PyIP, nr.Macs)
	return
}
