package main

import "fmt"

type salarycalculate interface {
	calculatesalary()
}
type leavestaken interface {
	calculateleaves() int
}
type employeeoperater interface {
	salarycalculate
	leavestaken
}

type employee struct {
	firstname   string
	lastname    string
	age         int
	salary      int
	pf          int
	totalleaves int
	leavestaken int
}

func main() {
	e := employee{
		firstname:   "devi",
		lastname:    "nallendhiran",
		age:         25,
		salary:      30000,
		pf:          2000,
		totalleaves: 30,
		leavestaken: 5,
	}

	var empop employeeoperater = e
	empop.calculatesalary()
	//var l leavestaken = e
	fmt.Println("leavestaken ", empop.calculateleaves())
}
func (e employee) calculatesalary() {
	fmt.Printf("%s %s salary is $%d\n", e.firstname, e.lastname, (e.salary + e.pf))

}
func (e employee) calculateleaves() int {
	return e.totalleaves - e.leavestaken

}
