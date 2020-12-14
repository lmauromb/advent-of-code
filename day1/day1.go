package day1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const goal = 2020

// Input shape of expenses' input
type Input struct {
	Expenses []int `json:"expenses"`
}

func loadExpenses(fileName string) ([]int, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return []int{}, err
	}

	var input Input

	err = json.Unmarshal(data, &input)
	if err != nil {
		return []int{}, err
	}

	return input.Expenses, nil
}

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			curr := arr[j]
			next := arr[j+1]

			if curr > next {
				arr[j] = next
				arr[j+1] = curr
			}
		}
	}

	return arr
}

func findPair(unsortedArr []int) (int, int) {
	sortedArr := bubbleSort(unsortedArr)

	var floorIndex = 0
	var topIndex = len(sortedArr) - 1

	for floorIndex < topIndex {
		pairSum := sortedArr[floorIndex] + sortedArr[topIndex]
		if pairSum == goal {
			return sortedArr[floorIndex], sortedArr[topIndex]
		}
		if pairSum < goal {
			floorIndex++
			continue
		}
		if pairSum > goal {
			topIndex--
			continue
		}
	}
	return 0, 0
}

func findTriplet(unsortedArr []int) (int, int, int) {
	sortedArr := bubbleSort(unsortedArr)

	for i := 0; i < len(sortedArr); i++ {
		var floorIndex = i + 1
		var topIndex = len(sortedArr) - 1

		for floorIndex < topIndex {
			tipletSum := sortedArr[floorIndex] + sortedArr[topIndex] + sortedArr[i]
			if tipletSum == goal {
				return sortedArr[floorIndex], sortedArr[topIndex], sortedArr[i]
			}
			if tipletSum < goal {
				floorIndex++
				continue
			}
			if tipletSum > goal {
				topIndex--
				continue
			}
		}
	}

	return 0, 0, 0
}

// DisplayExpenses prints out input file
func DisplayExpenses() {
	expns, err := loadExpenses("./day1/input.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	a, b := findPair(expns)

	fmt.Println("Day 1 Part I")
	fmt.Printf("Pair: %v, %v\n", a, b)
	fmt.Printf("Sum: %v\n", a+b)
	fmt.Printf("Product: %v\n", a*b)

	fmt.Println("Day 1 Part II")
	a, b, c := findTriplet(expns)
	fmt.Printf("Triplet: %v, %v, %v\n", a, b, c)
	fmt.Printf("Sum: %v\n", a+b+c)
	fmt.Printf("Product: %v\n", a*b*c)
}
