package main

import "fmt"

type salarycalculate interface {
	calculatesalary() int
}

type permanet struct {
	empId  int
	salary int
	pf     int
}
type contract struct {
	empId  int
	salary int
}

func main() {
	emp1 := permanet{
		empId:  1,
		salary: 15000,
		pf:     900,
	}
	emp2 := permanet{
		empId:  2,
		salary: 20000,
		pf:     1000,
	}
	emp3 := contract{
		empId:  3,
		salary: 35000,
	}
	emp4 := contract{
		empId:  4,
		salary: 30000,
	}
	employees := []salarycalculate{emp1, emp2, emp3, emp4}
	totalsalary(employees)
}
func (p permanet) calculatesalary() int {
	return p.salary + p.pf

}
func (c contract) calculatesalary() int {
	return c.salary

}
func totalsalary(s []salarycalculate) {
	expense := 0
	for _, v := range s {
		expense = expense + v.calculatesalary()

	}
	fmt.Printf("totalsalary is $%d", expense)
}
