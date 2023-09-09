package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Response struct {
	SlackName      string `json:"slack_name"`
	CurrentDay     string `json:"current_day"`
	CurrentUTCTime string `json:"utc_time"`
	Track          string `json:"track"`
	FileURL        string `json:"github_file_url"`
	SourceCodeURL  string `json:"github_repo_url"`
	StatusCode     int    `json:"status_code"`
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	slackName := "Adeyinka Boluwatife"
	Track := "backend"

	CurrentDay := time.Now().UTC().Format("Monday")
	currentTime := time.Now().UTC().Format("2006-01-02T15:04:05Z")
	GithubFileUrl := "https://github.com/mayowa1305/HNGx_UserInfo/blob/main/main.go"
	GithubUrl := "https://github.com/mayowa1305/HNGx_UserInfo"

	response := Response{
		SlackName:      slackName,
		CurrentDay:     CurrentDay,
		CurrentUTCTime: currentTime,
		Track:          Track,
		FileURL:        GithubFileUrl,
		SourceCodeURL:  GithubUrl,
		StatusCode:     http.StatusOK,
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func main() {
	http.HandleFunc("/get_info", getInfo)

	fmt.Println("server running on :8080")
	http.ListenAndServe(":8080", nil)
}
