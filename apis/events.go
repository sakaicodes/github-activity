package apis

import (
	"io"
	"net/http"
)

func GetEvents(username string) ([]byte, int) {
	req, err := http.NewRequest("GET", "https://api.github.com/users/"+username+"/events", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2026-03-10")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return body, resp.StatusCode

}
