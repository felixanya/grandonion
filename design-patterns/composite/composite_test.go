package composite

import "testing"

func TestNewDeveloper(t *testing.T) {
	manager1 := NewManager("aaron", "dept.1-manager", 3000)
	developer1 := NewDeveloper("jessie", "dept.1-developer", 1000)
	developer2 := NewDeveloper("shawn", "dept.1-developer", 1000)

	manager1.AddEmp(developer1)
	manager1.AddEmp(developer2)

	manager1.PrintInfo()

	developer3 := NewDeveloper("jake", "dept.1-developer", 1000)
	manager1.AddEmp(developer3)
	manager1.PrintInfo()

	manager2 := NewManager("jeff", "dept.12-manager", 2500)
	developer4 := NewDeveloper("chris", "dept.12-developer", 1000)
	manager2.AddEmp(developer4)
	manager1.AddEmp(manager2)

	manager1.PrintInfo()
	manager2.PrintInfo()
}
