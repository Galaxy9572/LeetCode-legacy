package main

func game(guess []int, answer []int) int {
	count := 0
	for index := range guess {
		if guess[index] == answer[index] {
			count++
		}
	}
	return count
}