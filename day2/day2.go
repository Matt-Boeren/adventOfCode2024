package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day2Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	amountOfSafe := 0
	for scanner.Scan() {
		lineInput := scanner.Text()
		splitInput := strings.Split(lineInput, " ")
		var report []int
		for i := 0; i < len(splitInput); i++ {

			num, err := strconv.Atoi(splitInput[i])
			if err != nil {
				fmt.Println("can't convert this to an int!")
			} else {
				report = append(report, num)
			}
		}
		result := checkReport(report)
		if result == true {
			amountOfSafe++
		} else {
			for i := 0; i < len(report); i++ {
				newReport := append([]int(nil), report...)
				newReport = append(newReport[:i], newReport[i+1:]...)
				newResult := checkReport(newReport)
				if newResult == true {
					amountOfSafe++
					break
				}
			}
		}
	}
	fmt.Println(amountOfSafe)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
func checkReport(l []int) bool {
	increase := false
	list := l
	safe := true
	for i := 0; i < len(list)-1; i++ {
		if i == 0 && list[i] < list[i+1] {
			increase = true
		}
		if increase == true {
			if list[i] >= list[i+1] {
				safe = false
				break
			} else {
				if list[i+1]-list[i] > 3 {
					safe = false
					break
				}
			}
		} else {
			if list[i] <= list[i+1] {
				safe = false
				break
			} else {
				if list[i]-list[i+1] > 3 {
					safe = false
					break
				}
			}
		}
	}
	return safe
}
