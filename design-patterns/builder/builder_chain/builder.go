package builder

import "fmt"

type IAllocator interface {
	ApplyIP() IAllocator
	GenMac() IAllocator
	AllocPyHost() IAllocator
	Done() (IResource, error)
}

type IResource interface {
	ListAllocated()
}

type Allocator struct {
	Allocated 	*NetResource
	err 		error
}

func NewAllocator() *Allocator {
	return &Allocator{
		Allocated: &NetResource{},
		err: nil,
	}
}

func (al *Allocator) ApplyIP() IAllocator {
	fmt.Println("applying ip address...")
	ip := "192.168.168.11"
	al.Allocated.IP = ip
	return al
}

func (al *Allocator) GenMac() IAllocator {
	fmt.Println("generating mac address...")
	al.Allocated.Macs = []string{"mad1", "mac2", "mac3"}
	return al
}

func (al *Allocator) AllocPyHost() IAllocator {
	fmt.Println("allocating physical host...")
	al.Allocated.PyIP = "192.168.168.10"
	al.Allocated.IpSeg = "192.168.168.0/24"
	return al
}

func (al *Allocator) Done() (IResource, error) {
	fmt.Println("Done!")
	if al.err != nil {
		return al.Allocated, al.err
	}
	return al.Allocated, nil
}

type NetResource struct {
	IP 		string
	PyIP 	string
	IpSeg 	string
	Macs 	[]string
}

func (nr *NetResource) ListAllocated() {
	fmt.Printf("Allocated resource:\n IP:\t%s\n PyIP:\t%s\n Macs:\t%s\n", nr.IP, nr.PyIP, nr.Macs)
	return
}
