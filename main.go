package main

import (
	"bufio"
	"os"
	"fmt"
)

type Log struct {
	eventType	string `json:"eventType"`
	message		string `json:"message"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter eventType: ")
	inputEventType, _ := reader.ReadString('\n')
	fmt.Print("Enter message: ")
	inputMessage, _ := reader.ReadString('\n')

	var currentLog Log;
	currentLog.eventType = inputEventType
	currentLog.message = inputMessage
	
	fmt.Print("Log: ", currentLog);
}