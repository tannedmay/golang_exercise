package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type IssueSeachResult struct {
	TotalCount int `json:"total_count"`
	Items      *[]Issue
}

type Issue struct {
	HTMLURL   string `json:"html_url"`
	Number    int
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

const IssueURL = "https://api.github.com/search/issues"

func searchIssues(queries []string) (*IssueSeachResult, error) {
	q := url.QueryEscape(strings.Join(queries, " "))
	resp, err := http.Get(IssueURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("searchIssue: http get failed with status: %s\n", resp.Status)
	}

	result := IssueSeachResult{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func main() {
	pastDay := make([]*Issue, 0)
	pastMonth := make([]*Issue, 0)
	pastYear := make([]*Issue, 0)
	pastMoreThanYear := make([]*Issue, 0)

	now := time.Now()
	result, err := searchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	for _, issue := range *result.Items {
		fmt.Printf("%s\n", issue.CreatedAt)
		switch {
		case issue.CreatedAt.After(now.AddDate(0, 0, -1)):
			pastDay = append(pastDay, &issue)
		case issue.CreatedAt.After(now.AddDate(0, -1, 0)):
			pastMonth = append(pastMonth, &issue)
		case issue.CreatedAt.After(now.AddDate(-1, 0, 0)):
			pastYear = append(pastYear, &issue)
		default:
			pastMoreThanYear = append(pastMoreThanYear, &issue)
		}
	}
	if len(pastDay) > 0 {
		fmt.Printf("Past day issues:\n")
		for _, iss := range pastDay {
			fmt.Printf("#%-5d %9.9s %.55s %v\n", iss.Number, iss.User.Login, iss.Title, iss.CreatedAt)
		}
	}
	if len(pastMonth) > 0 {
		fmt.Printf("Past month issues:\n")
		for _, iss := range pastMonth {
			fmt.Printf("#%-5d %9.9s %.55s %v\n", iss.Number, iss.User.Login, iss.Title, iss.CreatedAt)
		}
	}
	if len(pastYear) > 0 {
		fmt.Printf("Past year issues:\n")
		for _, iss := range pastYear {
			fmt.Printf("#%-5d %9.9s %.55s %v\n", iss.Number, iss.User.Login, iss.Title, iss.CreatedAt)
		}
	}
	if len(pastMoreThanYear) > 0 {
		fmt.Printf("Past more than a year issues:\n")
		for _, iss := range pastMoreThanYear {
			fmt.Printf("#%-5d %9.9s %.55s %v\n", iss.Number, iss.User.Login, iss.Title, iss.CreatedAt)
		}
	}
}
