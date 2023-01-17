package main

import "fmt"

var words = []string{"тест", "привет", "мир"}

func main() {
	// for each word:
	for _, word := range words {
		fmt.Println("GUESS THE WORD!\n")

		// create slace of runes based on the word
		wordRune := []rune(word)
		wordMap := make(map[rune]bool)

		for i := 0; i < len(wordRune); i++ {
			wordMap[wordRune[i]] = false
		}

		// for the new input
		var input string
		var finalWord string

		for {
                        fmt.Println("put a letter: ")
			fmt.Scanln(&input)

			inputRune := []rune(input)
			wordMap[inputRune[0]] = true

			// print the final result
			finalWord = ""
			for _, v := range wordRune {
				if wordMap[v] {
					finalWord = finalWord + string(v)
				} else {
					finalWord = finalWord + "*"
				}
			}
			fmt.Println(finalWord)

			res := 0
			for _, v := range wordRune {
				if wordMap[v] {
					res++
				}
			}
			if len(wordRune) == res {
				break
			}
		}
	}
}

