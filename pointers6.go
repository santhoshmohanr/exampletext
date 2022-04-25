package main

import "fmt"

type permanet struct {
	empId  int
	salary int
	pf     int
}
type contract struct {
	empId  int
	salary int
}
type freelancer struct {
	empId       int
	rateperhour int
	totalhours  int
}
type salarycalculate interface {
	calculatesalary() int
}

func main() {
	emp1 := permanet{
		empId:  1,
		salary: 20000,
		pf:     1000,
	}
	emp2 := permanet{
		empId:  2,
		salary: 25000,
		pf:     1200,
	}
	emp3 := contract{
		empId:  3,
		salary: 30000,
	}
	emp4 := contract{
		empId:  4,
		salary: 30000,
	}
	freelancer1 := freelancer{
		empId:       5,
		rateperhour: 70,
		totalhours:  120,
	}
	freelancer2 := freelancer{
		empId:       6,
		rateperhour: 100,
		totalhours:  100,
	}
	employees := []salarycalculate{emp1, emp2, emp3, emp4, freelancer1, freelancer2}
	totalsalary(employees)

}
func (p permanet) calculatesalary() int {
	return p.salary + p.pf

}
func (c contract) calculatesalary() int {
	return c.salary

}
func (f freelancer) calculatesalary() int {
	return f.rateperhour * f.totalhours

}
func totalsalary(s []salarycalculate) {
	expense := 0
	for _, v := range s {
		expense = expense + v.calculatesalary()
	}
	fmt.Printf("Totalsalary is $%d ", expense)
}
