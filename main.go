package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("----------------------------------------------")
	fmt.Println("===> Anagram <===")
	fmt.Println("----------------------------------------------")
	//define a new scanner to let user input the data
	var scanner = bufio.NewScanner(os.Stdin)

	//call getArrayLength to get the number of words from user
	arrayLen := getArrayLength(scanner)

	//call getArrayData to let user input the words
	arrayData := getArrayData(arrayLen, scanner)

	//call createAnagram function and get a map of sorted and original words
	sortedWords := createAnagram(arrayData)

	//print the result by looping through the values from the map
	var anagram [][]string
	fmt.Println("----------------------------------------------")
	fmt.Println("===> Anagram Result <===")
	fmt.Println("----------------------------------------------")
	for _, values := range sortedWords {
		anagram = append(anagram, values)
	}
	fmt.Println(anagram)
}

func getArrayData(arrayLen int, scanner *bufio.Scanner) []string {
	//define an empty array to store inputted data
	fmt.Println("----------------------------------------------")
	fmt.Println("Input the words")
	var arrayData []string

	//scan data inputted by the user and append it to the array
	for i := 1; i <= arrayLen; i++ {
		fmt.Printf("Data %d = ", i)
		scanner.Scan()
		stringData := scanner.Text()
		arrayData = append(arrayData, stringData)
	}
	return arrayData
}

func getArrayLength(scanner *bufio.Scanner) int {
	//loop infinitely until user inputted correct number
	for {
		//scan the length of array from user and convert it to an integer
		fmt.Println("----------------------------------------------")
		fmt.Print("Input number of words = ")
		scanner.Scan()
		arrayLen, err := strconv.Atoi(scanner.Text())

		//check for errors
		if err == nil && arrayLen >= 0 {
			return arrayLen
		}
		fmt.Println("Invalid number")
	}
}

func createAnagram(words []string) map[string][]string {
	//initialize a new map to store words with sorted characters as keys and the inputted words as values
	sortedWords := make(map[string][]string)

	//loop through every word from all inputed words
	for _, word := range words {
		//create a slice of runes from current looped word
		chars := []rune(word)

		//loop through every character in the current looped word
		charsLen := len(chars)
		for i := 0; i < charsLen-1; i++ {
			//loop through every characters after the current looped character in the current looped word
			for j := 0; j < charsLen-i-1; j++ {
				//sort the characters
				if chars[j] > chars[j+1] {
					chars[j], chars[j+1] = chars[j+1], chars[j]
				}
			}
		}
		//convert slice of runes back to a string
		sortedWord := string(chars)

		//append the current looped word to the initialized map
		sortedWords[sortedWord] = append(sortedWords[sortedWord], word)
	}
	return sortedWords
}
