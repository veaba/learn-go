package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

// func main() {
// 	client := github.NewClient(nil)

// 	// list all organizations for user "willnorris"
// 	orgs, _, err := client.Organizations.List(ctx, "veaba", nil)
// }

type CommitAuthor struct {
	Date  *time.Time `json:"date,omitempty"`
	Name  *string    `json:"name,omitempty"`
	Email *string    `json:"email,omitempty"`

	// The following fields are only populated by Webhook events.
	Login *string `json:"username,omitempty"` // Renamed for go-github consistency.
}
type RepositoryContentFileOptions struct {
	Message   *string       `json:"message,omitempty"`
	Content   []byte        `json:"content,omitempty"` // unencoded
	SHA       *string       `json:"sha,omitempty"`
	Branch    *string       `json:"branch,omitempty"`
	Author    *CommitAuthor `json:"author,omitempty"`
	Committer *CommitAuthor `json:"committer,omitempty"`
}

func main1() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "ttttttttttttt"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user

	message := "update(robot): docs/a.md"
	sha := "tttttttttttttttt"
	contentOps := &github.RepositoryContentFileOptions{
		Message:   &message,
		Content:   []byte("hello world wa"),
		Committer: &github.CommitAuthor{Name: github.String("veaba-bot"), Email: github.String("bot@veaba.me")},
		SHA:       &sha,
	}
	repos, _, err := client.Repositories.UpdateFile(ctx, "veaba", "vuepress-plugin-editable", "a.md", contentOps)
	// client.Repositories.GetCommitSHA1()
	fmt.Println("err=>", err)
	fmt.Println("repos=>", repos)

}
