package main

import (
	"fmt"
	"strconv"
)

func main() {
	var massMark, massJohn int
	var heightMark, heightJohn float32
	var bmiMark, bmiJohn float32
	var markHigherBmi bool
	
	// Menyimpan data massa dan tinggi Mark dan John.
	massMark = 78
	massJohn = 92
	heightMark = 1.69
	heightJohn = 1.95
	
	// Menghitung BMI Mark dan John
	bmiMark = countBMI(massMark, heightMark)
	bmiJohn = countBMI(massJohn, heightJohn)
	
	// Menyimpan informasi apakah BMI Mark lebih dari BMI John.
	markHigherBmi = bmiMark > bmiJohn
	
	fmt.Println("========== Data 1 ==========")
	printWithTemplate(massMark, heightMark, massJohn, heightJohn, bmiMark, bmiJohn, markHigherBmi)
	
	// Menyimpan data massa dan tinggi Mark dan John.
	massMark = 95
	massJohn = 85
	heightMark = 1.88
	heightJohn = 1.76
	
	// Menghitung BMI Mark dan John
	bmiMark = countBMI(massMark, heightMark)
	bmiJohn = countBMI(massJohn, heightJohn)
	
	// Menyimpan informasi apakah BMI Mark lebih dari BMI John.
	markHigherBmi = bmiMark > bmiJohn
	
	fmt.Println()
	fmt.Println("========== Data 2 ==========")
	printWithTemplate(massMark, heightMark, massJohn, heightJohn, bmiMark, bmiJohn, markHigherBmi)

}

func countBMI(mass int, height float32) float32 {
	return float32(mass) / (height*height)
}

/*
	Fungsi ini mencetak informasi massa, tinggi, dan BMI Mark dan John.
	Penulisan float dilakukan dengan presisi dua angka di belakang koma.
*/
func printWithTemplate(massMark int, heightMark float32, massJohn int, heightJohn float32, bmiMark float32, bmiJohn float32, markHigherBmi bool) {
	fmt.Printf("Mark's mass = %d kilograms\n", massMark)
	fmt.Printf("Mark's height = %.2f meters\n", heightMark)
	fmt.Printf("John's mass = %d kilograms\n", massJohn)
	fmt.Printf("John's height = %.2f meters\n", heightJohn)
	fmt.Println("---------- calculate BMI ----------")
	fmt.Printf("Mark's BMI = %.2f\n", bmiMark)
	fmt.Printf("John's BMI = %.2f\n", bmiJohn)
	fmt.Printf("Is Mark's BMI higher than John's? %s!!!\n", strconv.FormatBool(markHigherBmi))
}
	