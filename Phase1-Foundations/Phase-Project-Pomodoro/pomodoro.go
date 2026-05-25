package main

import (
	"fmt"
	"time"
)

type timer interface {
	getTime() int
}

type workSession struct {
	Time int
}
type shortBreak struct {
	Time int
}

type longBreak struct {
	Time int
}

func (w workSession) getTime() int {
	return w.Time
}

func (s shortBreak) getTime() int {
	return s.Time
}

func (l longBreak) getTime() int {
	return l.Time
}

// func updateActiveTimer(t timer) {
// }
//
// func counter(chan<- timer, <-chan int) {
// }

func addToCounter(adder chan<- int) {
	time.Sleep(time.Second * 1)
	adder <- 1
	// close(adder)
}

func timerCycle() {
	// w := workSession{Time: 25}
	// s := shortBreak{Time: 5}
	// l := longBreak{Time: 15}
}

//	func counter(actionInput <-chan string) {
//		action := <-actionInput
//		if action == "start" {
//			timeLeft := 25
//			ticker := time.NewTicker(1 * time.Second)
//			defer ticker.Stop()
//
//			fmt.Println("Starting 25-second countdown...")
//			for timeLeft > 0 {
//				select {
//				case <-ticker.C:
//					// \r clears the current line and returns the cursor to the start.
//					// This overwrites the previous number instead of printing a new line.
//					fmt.Printf("\rTime remaining: %2d seconds", timeLeft)
//					timeLeft--
//				}
//			}
//		}
//	}
// func getUseractions() {
// 	var action string
// 	ch := make(chan string)
// 	for {
// 		fmt.Printf("enter action: ")
//
// 		_, err := fmt.Scanf("%s", &action)
// 		if action == "start" {
// 			go counter(ch)
// 		}
// 		// fmt.Printf("%d number of items scanned: \n", input)
// 		// fmt.Printf("%s \n", action)
// 		if err != nil {
// 			fmt.Printf("")
// 		}
// 	}
// }

func main() {
	UserActions := make(chan string)
	timeCounter := make(chan int)
	action := "initial"
	currentTime := 25
	go func() {
		for {
			timeCounter <- 1
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for {
			var action string

			time.Sleep(1 * time.Millisecond)
			fmt.Printf("enter action: \n")

			_, err := fmt.Scanf("%s", &action)
			// if action == "start" {
			// 	UserActions <- "start"
			// }
			UserActions <- action
			// fmt.Printf("%d number of items scanned: \n", input)
			// fmt.Printf("%s \n", action)
			if err != nil {
				fmt.Printf("")
			}
		}
	}()
	// counter := make(chan int)
	// for range 25 {
	// 	go addToCounter(counter)
	// 	num := <-counter
	// 	fmt.Printf("%d \n", num)
	// }
	for {
		select {
		case item := <-UserActions:
			// Handle state changes and immediate commands here
			switch item {
			case "start":
				action = "start"

				fmt.Println("Timer started.")
			case "pause":
				action = "pause"
				fmt.Println("Timer paused.")
			case "restart":
				currentTime = 25
				action = "start" // Automatically resume counting after restart
				fmt.Println("Timer reset to 25.")
			case "show":
				// Print immediately. Do NOT overwrite the 'action' variable,
				// so if it was "start", it keeps running in the background.
				fmt.Printf("Current time left: %d seconds\n", currentTime)
			case "end":
				currentTime = 25
				action = "end"
				fmt.Printf("timer ended: \n")
			}

		case item := <-timeCounter:
			// This case ONLY handles the passage of time
			if action == "start" {
				currentTime = currentTime - item

				if currentTime <= 0 {
					fmt.Println("\nTime's up! Take a break.")
					action = "stopped"
				}
			}
		}
	}
}
