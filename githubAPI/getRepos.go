package main

import (
	"log"
	"os"

	"github.com/levigross/grequests"
)

// GithubToken is the github token added to my box
var GithubToken = os.Getenv("GITHUB_TOKEN")
var requestOptions = &grequests.RequestOptions{
	Auth: []string{GithubToken, "x-oauth-basic"}}

// Repo is the struct for the github repos
type Repo struct {
	ID       int    ` json:"id"`
	Name     string ` json:"name"`
	FullName string ` json:"full_name"`
	Forks    int    ` json:"forks"`
	Private  bool   ` json:"private"`
}

func getStats(url string) *grequests.Response {
	resp, err := grequests.Get(url, requestOptions)

	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	return resp
}

func main() {
	var repos []Repo
	var repoURL = "https://api.github.com/users/torvalds/repos"
	resp := getStats(repoURL)
	resp.JSON(&repos)
	log.Println(repos)
}
