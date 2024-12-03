package main

import(
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func main (){
	file, err :=  os.Open("day2Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	amountOfSafe := 0
	for scanner.Scan() {
		lineInput := scanner.Text()

		result := strings.Split(lineInput, " ")
		var list []int
		for i := 0; i < len(result); i++{
			
			num, err := strconv.Atoi(result[i])
			if err != nil{
				fmt.Println("can't convert this to an int!")
			}else{
				list = append(list, num)
			}
		}
		increase := false
		safe := true
		for i := 0; i < len(list) - 1; i++{
			if i == 0 && list[i] < list[i+1]{
				increase = true
			}
			if increase == true{
				if list[i] >= list[i+1]{
					safe = false
					break
				}else{
					if list[i+1] - list[i] > 3{
						safe = false
						break
					}
				}
			}else{
				if list[i] <= list[i+1]{
					safe = false
					break
				}else{
					if list[i] - list[i+1] > 3{
						safe = false
						break
					}
				}
			}
		}
		if safe == true{
			amountOfSafe++
		}
	}
	fmt.Println(amountOfSafe)
	if err := scanner.Err(); err != nil{
		log.Fatal(err)
	}
}
