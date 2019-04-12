package builder

import "testing"

func TestNewOrchestrate(t *testing.T) {
	allocator := &Allocator{
		netRes:NetResource{},
	}
	orch := NewOrchestrate(allocator)
	orch.Orchestrate()
	allocator.netRes.ListAllocated()
}
