package main

import "fmt"

func main() {
 
 score := 85

 if score >= 90 {
  fmt.Println("You got an A!")
 } else if score >= 80 {
  fmt.Println("You got a B.")
 } else if score >= 70 {
  fmt.Println("You got a C.")
 } else if score >= 60 {
  fmt.Println("You got a D.")
 } else {
  fmt.Println("You failed the exam.")
 }


}