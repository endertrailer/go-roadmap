package main

import (
	"fmt"
	"time"
)

type timer interface {
	getTime() int
	getTitle() string
}

type workSession struct {
	Time  int
	Title string
}

type shortBreak struct {
	Time  int
	Title string
}

type longBreak struct {
	Time  int
	Title string
}

func (w workSession) getTime() int     { return w.Time }
func (s shortBreak) getTime() int      { return s.Time }
func (l longBreak) getTime() int       { return l.Time }
func (w workSession) getTitle() string { return w.Title }
func (s shortBreak) getTitle() string  { return s.Title }
func (l longBreak) getTitle() string   { return l.Title }

func changeSession(sessionCount int) (int, timer) {
	sessionTypes := []timer{
		workSession{Time: 25, Title: "Work Session"},
		shortBreak{Time: 5, Title: "Short break"},
		longBreak{Time: 15, Title: "Long Break"},
	}

	if sessionCount >= 7 {
		return 0, sessionTypes[0]
	}

	nextSession := sessionCount + 1

	if nextSession == 7 {
		return nextSession, sessionTypes[2]
	}

	if nextSession%2 != 0 {
		return nextSession, sessionTypes[1]
	}

	return nextSession, sessionTypes[0]
}

func main() {
	UserActions := make(chan string)
	timeCounter := make(chan int)
	action := "initial"
	sessionCount := 0
	currentTime := 25 * 60

	var activeSession timer = workSession{Time: 25, Title: "Work Session"}

	go func() {
		for {
			timeCounter <- 1
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			var actionInput string
			time.Sleep(1 * time.Millisecond)
			fmt.Printf("enter action: \n")

			_, err := fmt.Scanf("%s", &actionInput)
			if err == nil {
				UserActions <- actionInput
			}
		}
	}()

	for {
		select {
		case item := <-UserActions:
			switch item {
			case "start":
				action = "start"
				fmt.Println("Timer started.")
			case "pause":
				action = "pause"
				fmt.Println("Timer paused.")
			case "restart":
				currentTime = activeSession.getTime() * 60
				action = "start"
				fmt.Printf("Timer reset for %s \n", activeSession.getTitle())
			case "show":
				mins := currentTime / 60
				secs := currentTime % 60
				fmt.Printf("Current time left: %02d:%02d\n", mins, secs)
			case "next":
				sessionCount, activeSession = changeSession(sessionCount)
				currentTime = activeSession.getTime()
				action = "start"
				fmt.Printf("moving to: %s \n", activeSession.getTitle())
			}

		case item := <-timeCounter:
			if action == "start" {
				currentTime = currentTime - item

				if currentTime <= 0 {
					sessionCount, activeSession = changeSession(sessionCount)
					currentTime = activeSession.getTime()
					action = "start"
					fmt.Printf("moving to: %s \n", activeSession.getTitle())
				}
			}
		}
	}
}
