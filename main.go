package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	droneURL := "https://drone.nflex.io"
	gogsURL := os.Getenv("DRONE_NETRC_MACHINE")
	droneUser := os.Getenv("DRONE_NETRC_USERNAME")
	dronePass := os.Getenv("DRONE_NETRC_PASSWORD")
	owner := os.Getenv("DRONE_REPO_OWNER")
	repo := os.Getenv("DRONE_REPO_NAME")
	buildNumber := os.Getenv("DRONE_BUILD_NUMBER")
	buildStatus := os.Getenv("CI_BUILD_STATUS")
	index := os.Getenv("DRONE_PULL_REQUEST")

	message := buildMessage(droneURL, owner, repo, buildStatus, buildNumber)
	payload, err := json.Marshal(map[string]string{"body": message})
	if err != nil {
		fmt.Printf("Failed to prepare payload: %s", err)
		os.Exit(1)
	}
	url := fmt.Sprintf("https://%s/api/v1/repos/%s/%s/issues/%s/comments", gogsURL, owner, repo, index)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Printf("Failed to prepare request: %s", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(droneUser, dronePass)

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Failed to create issue comment: %s", err)
		os.Exit(1)
	}
}

func buildMessage(droneURL, owner, repo, buildStatus, buildNumber string) string {
	icon := fmt.Sprintf("%s/static/favicon.ico", droneURL)
	buildURL := fmt.Sprintf("%s/%s/%s/%s", droneURL, owner, repo, buildNumber)
	message := fmt.Sprintf("### ![](%s) Drone build #[%s](%s): %s", icon, buildNumber, buildURL, buildStatus)
	return message
}
