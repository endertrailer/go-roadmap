# Pomodoro Timer

A simple, interactive command-line Pomodoro Timer written in Go. 

## Overview

This application helps you manage your time using the Pomodoro Technique. It rotates between different phases:
- **Work Session**: 25 units
- **Short Break**: 5 units
- **Long Break**: 15 units

The standard cycle completes 4 work sessions (with short breaks in between) followed by a long break, before resetting the cycle.

## How to Run

Ensure you have Go installed on your system. Then, navigate to the project directory and run:

```bash
go run pomodoro.go
```

## Commands

While the timer is running (or paused), you can type the following commands into the terminal and press `Enter` to control it:

- `start`: Starts or resumes the timer.
- `pause`: Pauses the current timer.
- `restart`: Resets the timer for the current session and starts it again.
- `show`: Displays the remaining time for the current session.
- `next`: Skips the current session and immediately moves to the next one (e.g., skips a Work Session and goes straight to a Short Break).

## Example Usage

```text
enter action: 
start
Timer started.
enter action: 
show
Current time left: 21 seconds
enter action: 
pause
Timer paused.
```
