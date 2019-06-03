package builder

import (
	"fmt"
	"testing"
)

func TestNewAllocator(t *testing.T) {
	allocator := NewAllocator()
	allocated, err := allocator.AllocPyHost().ApplyIP().GenMac().Done()
	if err != nil {
		fmt.Println("分配资源错误,", err.Error())
		return
	}
	fmt.Println("--------------------")
	allocated.ListAllocated()
}
