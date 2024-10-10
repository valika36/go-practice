package main

import "fmt"

type FullTime struct {
	amount   int
	noOfDays int
}

type Contractor struct {
	amount   int
	noOfDays int
}

type FreeLancer struct {
	amount    int
	noOfHours int
}

type Salary interface {
	CalculateSalary() int
}

func (f FullTime) CalculateSalary() int {
	return f.noOfDays * f.amount
}

func (c Contractor) CalculateSalary() int {
	return c.noOfDays * c.amount
}

func (fl FreeLancer) CalculateSalary() int {
	return fl.noOfHours * fl.amount
}

func main() {
	fullTime := FullTime{
		amount:   15000,
		noOfDays: 30,
	}

	contractor := Contractor{
		amount:   3000,
		noOfDays: 30,
	}

	freelancer := FreeLancer{
		amount:    2000,
		noOfHours: 20,
	}

	fmt.Println("FullTime salary: ", fullTime.CalculateSalary())
	fmt.Println("Contractor salary: ", contractor.CalculateSalary())
	fmt.Println("FreeLancer salary: ", freelancer.CalculateSalary())
}
