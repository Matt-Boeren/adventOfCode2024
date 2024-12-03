package main

import(
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"sort"
	"math"
)

func main (){
	file, err :=  os.Open("day1Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var list1 []int
	var list2 []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineInput := scanner.Text()
		result := strings.Split(lineInput, "   ")
		l1, err1 := strconv.Atoi(result[0])
		if err1 != nil{
			fmt.Println("can't convert this to an int!")
		}else{
			list1 = append(list1, l1)
		}
		l2, err2 := strconv.Atoi(result[1])
		if err2 != nil{
			fmt.Println("cant convert this to an int!")
		}else{
			list2 = append(list2, l2)
		}
	}
	if err := scanner.Err(); err != nil{
		log.Fatal(err)
	}

	sort.Sort(sort.IntSlice(list1))
	sort.Sort(sort.IntSlice(list2))

	sum := 0
	for i := 0; i< len(list1); i++{
		dif := int(math.Abs(float64(list1[i] - list2[i])))
		sum += dif
		fmt.Println("l1 " + fmt.Sprint(list1[i]) + " l2 " + fmt.Sprint(list2[i]) + " dif " + fmt.Sprint(dif) + " sum " + fmt.Sprint(sum))
	}
	fmt.Println("sum: " + fmt.Sprint(sum))
}

