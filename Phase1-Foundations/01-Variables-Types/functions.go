package main

import "fmt"

func dfa(word string) string {
	activeState := 0
	byteString := []byte(word)
	for len(byteString) > 0 {
		switch activeState {
		case 0:
			q0(&byteString, &activeState)
		case 1:
			q1(&byteString, &activeState)
		case 2:
			q2(&byteString, &activeState)
		case 3:
			q3(&byteString, &activeState)
		default:
		}
	}
	if activeState == 3 {
		return "accepted"
	}
	return "failed"
}

func q0(currentWord *[]byte, activeState *int) {
	fmt.Printf("%s \n", currentWord)
	if (*currentWord)[0] == 'a' {
		*activeState = 1
		*currentWord = (*currentWord)[1:]
	} else {
		*currentWord = (*currentWord)[1:]
	}
}

func q1(currentWord *[]byte, activeState *int) {
	if (*currentWord)[0] == 'b' {
		*activeState = 2
		*currentWord = (*currentWord)[1:]
	} else {
		*currentWord = (*currentWord)[1:]
	}
}

func q2(currentWord *[]byte, activeState *int) {
	if (*currentWord)[0] == 'b' {
		*activeState = 3
		*currentWord = (*currentWord)[1:]
	} else {
		*activeState = 1
		*currentWord = (*currentWord)[1:]
	}
}

func q3(currentWord *[]byte, activeState *int) {
	if (*currentWord)[0] == 'b' {
		*activeState = 0
		*currentWord = (*currentWord)[1:]
	} else {
		*activeState = 1
		*currentWord = (*currentWord)[1:]
	}
}
