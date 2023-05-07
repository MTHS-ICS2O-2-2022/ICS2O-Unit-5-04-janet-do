// Copyright (c) 2020 Janet Do All rights reserved
//
// Created by: Janet Do
// Created on: Sep 2020
// This file checks the user's age to see what movies they can see

"use strict"

function calculate() {
  const age = document.getElementById("age").value
  document.getElementById("age").innerHTML = age
  const day = document.getElementById("day").value
  document.getElementById("day").innerHTML = day

  if (day === "tuesday" || day === "thursday" || (age >= 12 && age <= 21)) {
    document.getElementById("determine").innerHTML =
      "You can get into the musuem."
  } else {
    document.getElementById("determine").innerHTML =
      "Sorry you can't get into the musuem"
  }
}
