package main

import "fmt"

type Employee interface {
	CalculateSalary() int
}

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

func CalculateSalary(e Employee) int {
	return e.CalculateSalary()
}

func (c FullTime) CalculateSalary() int {
	return c.noOfDays * c.amount
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

	fmt.Println("FullTime salary: ", CalculateSalary(fullTime))
	fmt.Println("Contractor salary: ", CalculateSalary(contractor))
	fmt.Println("FreeLancer salary: ", CalculateSalary(freelancer))
}
