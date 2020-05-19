package main

import (
	"fmt"
	"flag"
)

const (
	py = 13.0
	goGrade = 16.0
	general = 11.0
	frontend = 22.0
	cSharp = 16.0
)

var python, golang, gen, fend, csharp bool
func init() {
	const (
		defaultOn = true
		defaultOff = false
		usage = "Sets the courses a student is offering. The defaults are General and Go"
	)
	//set flags
	flag.BoolVar(&python,"python",defaultOff,"Includes python as one of your course")
	flag.BoolVar(&python,"p",defaultOff,"Same as -python")

	flag.BoolVar(&gen,"general",defaultOn,"Includes general as one of your course.\n")
	flag.BoolVar(&gen,"gen",defaultOn,"Same as -general")

	flag.BoolVar(&golang,"golang",defaultOff,"Includes golang as one of your course\n")
	flag.BoolVar(&golang,"g",defaultOff,"Same as -golang")
	
	
	flag.BoolVar(&fend,"frontend",defaultOff,"Includes frontend as one of your course\n")
	flag.BoolVar(&fend,"f",defaultOff,"Same as -frontend")
	
	flag.BoolVar(&csharp,"csharp",defaultOff,"Includes C# as one of your course\n")
	flag.BoolVar(&csharp,"c",defaultOff,"Same as -csharp")
	
	flag.Parse()  // Execute command-line parsing of flags
}

func goScore() float64{
	var task1, task2, task3, task4 float64
	fmt.Println("Enter scores for go task")
	fmt.Scanf("%f %f %f %f", &task1, &task2, &task3, &task4)
	total := task1 + task2 + task3 + task4
	return total
}

func generalScore() float64{
	var task1, task2, task3, task4, task5 float64
	fmt.Println("Enter scores for general tasks")
	fmt.Scanf("%f %f %f %f %f", &task1, &task2, &task3, &task4, &task5)
	fmt.Println(task1, task2, task3, task4, task5)
	total := task1 + task2 + task3 + task4 + task5
	return total
}

func pyScore() float64 {
	var task1, task2, task3, task4 float64
	fmt.Println("Enter scores for python tasks")
	fmt.Scanf("%f %f %f %f", &task1, &task2, &task3, &task4)
	fmt.Println(task1, task2, task3, task4)
	total := task1 + task2 + task3 + task4
	return total
}

func frontendScore() float64{
	var task1, task2, task3, task4 float64
	fmt.Println("Enter scores for python tasks")
	fmt.Scanf("%f %f %f %f", &task1, &task2, &task3, &task4)
	total := task1 + task2 + task3 + task4
	return total
}

func cSharpScore() float64 {
	var task1, task2, task3, task4 float64
	fmt.Println("Enter scores for python tasks")
	fmt.Scanf("%f %f %f %f", &task1, &task2, &task3, &task4)
	total := task1 + task2 + task3 + task4
	return total
}

func main() {
	var excellentScore, yourScore float64
	// accumulating your scores
	fmt.Println(python, golang, gen, fend, csharp)
	if gen{
		total := generalScore()
		excellentScore+=general
		yourScore+=total
	}
	
	if golang{
		total := goScore()
		excellentScore+=goGrade
		yourScore+=total
	}

	
	if fend{
		total := frontendScore()
		excellentScore+=frontend
		yourScore+=total
	}

	if csharp{
		total := cSharpScore()
		excellentScore+=cSharp
		yourScore+=total
	}

	
	if python{
		total := pyScore()
		excellentScore+=py
		yourScore+=total
		fmt.Println(excellentScore,yourScore)
	}


	percentageScore := (yourScore/excellentScore) *100
	if percentageScore >= (0.7*excellentScore){
		fmt.Printf("Hurray! You are a finalist with a score of %.2f",percentageScore)
	}else{
		fmt.Printf("Growth is incremental. Don't Stop!. Your grade: %.2f\n", percentageScore)
		fmt.Printf("You needed at least: %.2f", 0.7*excellentScore)
	}
	
}