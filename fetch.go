package main

import (
	"fmt"
	"time"
	"errors"
	"github.com/google/go-github/v45/github"
	"context"
	"strings"
)

func getTagDate(owner string, repo string, client *github.Client, sha string) (time.Time, error) {

	commit, resp, err := client.Repositories.GetCommit(context.Background(), owner, repo, sha, nil)

	if err != nil {
		return time.Time{}, err
	}

	if resp.StatusCode != 200 {
		fmt.Printf("Failed to fetch %s/%s | status code = %d\n", owner, repo, resp.StatusCode)
		return time.Time{}, errors.New("failed to get commit")
	} 

	return *commit.Commit.Committer.Date, nil
}

func ResolveReleaseDate(pName string, ver string) (time.Time, error) {

	//https://go.dev/blog/v2-go-modules
	// v2 postfix issue
	owner := strings.Split(pName, "/")[1]
	repo := strings.Split(pName, "/")[2]

	fmt.Printf("fetching... owner = %s | repo = %s\n", owner, repo)

	// if u need more rate limit read https://github.com/google/go-github
	client := github.NewClient(nil)

	tags, resp, err := client.Repositories.ListTags(context.Background(), owner, repo, nil)

	if err != nil {
		return time.Time{}, err
	}

	if resp.StatusCode != 200 || len(tags) == 0 {
		fmt.Printf("Failed to fetch %s/%s | status code = %d\n", owner, repo, resp.StatusCode)
		return time.Time{}, err
	} else {
		for _, tag := range tags {
			if *tag.Name == ver {
				fmt.Printf("Found tag %s\n", *tag.Name)
				
				return getTagDate(owner, repo, client, *tag.Commit.SHA)
			}
		}
	}

	return time.Time{}, nil
}