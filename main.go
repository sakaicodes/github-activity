package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/sakaicodes/github-activity/apis"
)

func main() {
	if len(os.Args) < 2 {
		println("Usage: github-activity <github-username>")
		os.Exit(1)
	}

	fmt.Println("Fetching GitHub activity for user:", os.Args[1])
	body, status := apis.GetEvents(os.Args[1])
	fmt.Println("Status Code:", status)

	var pretty bytes.Buffer
	json.Indent(&pretty, body, "", "  ")
	fmt.Println("Response Body:", pretty.String())
}
