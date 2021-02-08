package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

func main1() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "ddddddddddddd"},
	)
	tc := oauth2.NewClient(ctx, ts)

	fmt.Println("ctx=>", ctx)
	fmt.Println("ts=>", ts)
	fmt.Println("tc=>", tc)
	client := github.NewClient(tc)

	// list all repositories for the authenticated user

	// contentOps := &github.RepositoryContentGetOptions{
	// 	Ref: "main",
	// }
	repos, a, err := client.Users.Get(ctx, "yexiaoking")
	// client.Repositories.GetCommitSHA1()
	fmt.Println("repos=>", repos)
	fmt.Println("a=>", a)
	fmt.Println("err=>", err)
	fmt.Println("name=>", repos.GetLogin())
	fmt.Println("email=>", repos.GetEmail())

}
