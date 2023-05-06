// Copyright (c) 2020 Janet Do All rights reserved
//
// Created by: Janet Do
// Created on: Sep 2020
// This file checks the user's age to see what movies they can see
package main

import "fmt"

func main() {
	var age int
	var day string

	fmt.Println("Please enter your age:")
	fmt.Scanln(&age)

	fmt.Println("Please enter the day of the week:")
	fmt.Scanln(&day)

	if (day == "Tuesday" || day == "Thursday") && age >= 12 && age <= 21 {
		fmt.Println("\nCongratulations! You are eligible for discounted pricing.")
	} else {
		fmt.Println("\nSorry, you are not eligible for discounted pricing.")
	}
}


