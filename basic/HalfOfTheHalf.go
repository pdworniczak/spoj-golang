package main

import "fmt"

func main() {
	var wordsNumber int

	fmt.Scan(&wordsNumber)

	words := make([]string, wordsNumber)

	for i := 0; i < wordsNumber; i++ {
		fmt.Scan(&words[i])
	}

	for i := 0; i < wordsNumber; i++ {
		newWord := ""
		for i, v := range words[i][:len(words[i])/2] {
			if i%2 == 0 {
				newWord = newWord + string(v)
			}
		}

		fmt.Println(newWord)
	}
}
