package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

func main3() {
	createRR()
	// getRefList()
	// createRef()
	// getDefaultBranch()

}

// func getDefaultBranch() {
// 	ctx := context.Background()
// 	ts := oauth2.StaticTokenSource(
// 		&oauth2.Token{AccessToken: "rrrrrrrrrrrrrrr"},
// 	)
// 	tc := oauth2.NewClient(ctx, ts)
// 	client := github.NewClient(tc)

// 	// ref := github.Reference{
// 	// 	Ref:github.String("veaba-bot@123")
// 	// }
// 	repos, a, err := client.Repositories.Get(ctx, "veaba", "deno-docs")
// 	// client.Repositories.GetCommitSHA1()
// 	fmt.Println("repos=>", repos)
// 	fmt.Println("default branch=>", repos.GetDefaultBranch())
// 	fmt.Println("a=>", a)
// 	fmt.Println("err=>", err)
// }

func createRR() {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "vvvvvvvvvvvvv"},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	newBranch := "veaba-bot@r8l16jtn"
	updatePath := "a.md "
	newPullRequestOps := github.NewPullRequest{
		Title: github.String("Docs(veaba-bot): update " + updatePath + newBranch),
		Body:  github.String("The Pull Request Powered by @veaba-bot."),
		// Head: github.String("veaba:" + newBranch), // to
		Head: github.String(newBranch), // to
		// Base: github.String("veaba:master"),       // from
		Base: github.String("master"), // from
	}
	repos, _, err := client.PullRequests.Create(ctx, "veaba", "deno-docs", &newPullRequestOps)
	fmt.Println("repos=>", repos)
	fmt.Println("err=>", err)
	// fmt.Println("name=>", repos.GetLogin())
	// fmt.Println("email=>", repos.GetEmail())}
}
