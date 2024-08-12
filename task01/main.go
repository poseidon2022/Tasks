package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func gradeCalculator(subjects map[string]float64) float64 {
	var _sum float64
	for _, value := range subjects {
		_sum += value
	}

	var totalSubs float64 = float64(len(subjects))
	return _sum / totalSubs
}

func main() {
	//strconv.Atoi to convert the strings to an int
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter ur name here: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading the name!")
		return
	}

	input = strings.TrimSpace(input)
	fmt.Printf("Alright %v, how many subjects have u taken this semester?: ", input)
	subject, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading the number of subjects!")
		return 
	}

	subject = strings.TrimSpace(subject)
	subjectNum, err := strconv.Atoi(subject)
	for err != nil  {
		fmt.Println("Please enter a valid subject number!")
		fmt.Printf("Alright %v, how many subjects have u taken this semester?: ", input)
		subject, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading the number of subjects!")
			return 
		}
		subject = strings.TrimSpace(subject)
		subjectNum, err = strconv.Atoi(subject)
	}
	
	fmt.Printf("You have taken %v subjects\n", subjectNum)
	fmt.Println("-------------Please Enter the subject names along with your respective grades-----------")


	grades := map[string]float64{}
	for i := 0; i < subjectNum; i++ {
		fmt.Printf("Subject %v: ",i + 1)
		subjectName, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading the subject name!")
			return
		}

		subjectName = strings.TrimSpace(subjectName)
		fmt.Print("What was your grade for this subject: ")
		grade, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading the grade!")
			return
		}

		grade = strings.TrimSpace(grade)
		gradeNum, err := strconv.ParseFloat(grade, 64)
		for err != nil || ((gradeNum < 0) || (gradeNum > 100)) {
			fmt.Println("Please enter a valid input!")
			fmt.Print("What was your grade for this subject: ")
			grade, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading the grade!")
				return
			}

			grade = strings.TrimSpace(grade)
			gradeNum, err = strconv.ParseFloat(grade, 64)
		}
		grades[subjectName] = gradeNum
	}

	averageGrade := gradeCalculator(grades)
	fmt.Println(input)
	fmt.Println("------Subject Scores and Average------")

	for index, value := range grades {
		fmt.Printf("%v: %v\n", index, value)
	}

	fmt.Printf("Avergae Score: %v", averageGrade)

}