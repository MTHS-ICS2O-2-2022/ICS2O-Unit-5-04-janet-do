// Copyright (c) 2020 Janet Do All rights reserved
//
// Created by: Janet Do
// Created on: Sep 2020
// This file checks the user's age and the day of the week for admission
package main

import "fmt"

func main() {
	var age int
	var day string

	fmt.Println("Please enter your age:")
	fmt.Scanln(&age)

	fmt.Println("Please enter the day of the week:")
	fmt.Scanln(&day)

	if (day == "tuesday" || day == "thursday") && age >= 12 && age <= 21 {
		fmt.Println("\nCongratulations! You can get in")
	} else {
		fmt.Println("\nSorry, you can't get in")
	}
}
