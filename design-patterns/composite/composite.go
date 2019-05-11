package composite

import (
	"fmt"
)

// 组合模式，组合的是对象

// 通用employee接口
type IEmployee interface {
	PrintInfo()
}

// 角色接口
type IManager interface {
	IEmployee
	AddEmp(IEmployee)
	RemoveEmp(IEmployee)
}

type IDeveloper interface {
	IEmployee
	DevelopProject()
}

// manager角色类
type Manager struct {
	name      string
	salary    float32
	position  string
	Employees []IEmployee
}

func NewManager(name, position string, salary float32) IManager {
	return &Manager{
		name:      name,
		position:  position,
		salary:    salary,
		Employees: make([]IEmployee, 0),
	}
}

func (m *Manager) AddEmp(emp IEmployee) {
	m.Employees = append(m.Employees, emp)
}

func (m *Manager) RemoveEmp(emp IEmployee) {
	for i, e := range m.Employees {
		if e == emp {
			m.Employees = append(m.Employees[:i], m.Employees[i+1:]...)
		}
	}
}

func (m *Manager) PrintInfo() {
	fmt.Printf("Name: %s, Position: %s, Salary: %f, \nEmployees: \n", m.name, m.position, m.salary)
	for _, emp := range m.Employees {
		fmt.Printf("\t[%+v]\n", emp)
	}
}

// developer角色类
type Developer struct {
	name     string
	salary   float32
	position string
}

func NewDeveloper(name, position string, salary float32) IDeveloper {
	return &Developer{
		name:     name,
		position: position,
		salary:   salary,
	}
}

func (d *Developer) DevelopProject() {
	fmt.Println("some project developing")
}

func (d *Developer) PrintInfo() {
	fmt.Printf("Name: %s, Position: %s, Salary: %f", d.name, d.position, d.salary)
}
