package apis

import (
	"io"
	"net/http"
)

func GetEvents(username string) ([]byte, int) {
	req, err := http.Get("https://api.github.com/users/" + username + "/events")
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2026-03-10")

	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	return body, req.StatusCode

}
