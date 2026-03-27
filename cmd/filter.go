package cmd

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/sakaicodes/github-activity/apis"
)

func FilterEvent() {
	eventType := flag.String("filter", "", "Type of GitHub event to filter by (e.g., PushEvent, PullRequestEvent)")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		println("Usage: github-activity filter --username <github-username> [--event-type <event-type>]")
		os.Exit(1)
	}
	username := args[0]
	fmt.Println("Fetching events for user:", username)
	if *eventType == "" {
		fmt.Println("Filtering by event type: None (showing all events)")
	} else {
		fmt.Println("Filtering by event type:", *eventType)
	}
	fmt.Println()

	//API call to get events by type
	body, status := apis.GetEventsByType(username, *eventType)

	// Output the results
	fmt.Println("Status Code:", status)
	var pretty bytes.Buffer
	json.Indent(&pretty, body, "", "  ")
	fmt.Println("Response Body:", pretty.String())
}
