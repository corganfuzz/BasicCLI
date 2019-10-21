package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/levigross/grequests"
	"github.com/urfave/cli"
)

// GithubToken comment
var GithubToken = os.Getenv("GITHUB_TOKEN")
var requestOptions = &grequests.RequestOptions{
	Auth: []string{GithubToken, "x-oauth-basic"}}

// Repo Struct
type Repo struct {
	ID       int    ` json:"id"`
	Name     string ` json:"name"`
	FullName string ` json:"full_name"`
	Forks    int    ` json:"forks"`
	Private  bool   ` json:"private"`
}

// File struct
type File struct {
	Content string ` json:"content"`
}

// Gist struct
type Gist struct {
	Description string          ` json:"description"`
	Public      bool            ` json:"public"`
	Files       map[string]File ` json:"files"`
}

func getStats(url string) *grequests.Response {
	resp, err := grequests.Get(url, requestOptions)
	if err != nil {
		log.Fatalln("Can not make request", err)
	}
	return resp
}

func createGist(url string, args []string) *grequests.Response {

	description := args[0]

	var fileContents = make(map[string]File)

	for i := 1; i < len(args); i++ { //loop here because arguments are file names with paths

		dat, err := ioutil.ReadFile(args[i])

		if err != nil {
			log.Println("Check filenames. absolute path or same directorry are allowed")
			return nil
		}
		var file File
		file.Content = string(dat)
		fileContents[args[i]] = file
	}

	var gist = Gist{Description: description, Public: true, Files: fileContents}
	var postBody, _ = json.Marshal(gist)
	var requestOptionsCopy = requestOptions

	requestOptionsCopy.JSON = string(postBody) // Add data to JSON field

	resp, err := grequests.Post(url, requestOptionsCopy) // Make a post request to github

	if err != nil {
		log.Println("Create request failed for Github API")
	}
	return resp
}

func main() {

	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:    "fetch",
			Aliases: []string{"f"},
			Usage:   "Fetch the repo details with user. [Usage]: goTool fetch user",
			Action: func(c *cli.Context) error {
				if c.NArg() > 0 {
					var repos []Repo
					user := c.Args()[0]
					var repoURL = fmt.Sprintf("https://api.github.com/users/%s/repos", user)
					resp := getStats(repoURL)
					resp.JSON(&repos)
					log.Println(repos)
				} else {
					log.Println("Give a username. See -h to see help")
				}
				return nil
			},
		},
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "Creates a gist from the given text. [Usage]: goTool name 'description' sample.txt",
			Action: func(c *cli.Context) error {
				if c.NArg() > 1 {
					args := c.Args()
					var postURL = "https://api.github.com/gists"
					resp := createGist(postURL, args)
					log.Println(resp.String())
				} else {
					log.Println("Please give sufficient arguments. See -h to see help")
				}
				return nil
			},
		},
	}
	app.Version = "1.0"
	app.Run(os.Args)
}
