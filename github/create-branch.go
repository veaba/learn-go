package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

func main1() {
	getRef()
	// getRefList()
	// createRef()
	// getDefaultBranch()

}

// func getDefaultBranch() {
// 	ctx := context.Background()
// 	ts := oauth2.StaticTokenSource(
// 		&oauth2.Token{AccessToken: "xxx"},
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

func getRef() {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "bbbbbbbbbbbbbbb"},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	repos, _, err := client.Git.GetRef(ctx, "veaba", "deno-docs", "refs/heads/master")
	fmt.Println("repos=>", repos)
	fmt.Println("branch sha=>", repos.Object.GetSHA())
	fmt.Println("err=>", err)
	// fmt.Println("name=>", repos.GetLogin())
	// fmt.Println("email=>", repos.GetEmail())}
	/*
		github.Reference{Ref:"refs/heads/master", URL:"https://api.github.com/repos/veaba/deno-docs/git/refs/heads/master", Object:github.GitObject{Type:"commit", SHA:"3333", URL:"https://api.github.com/repos/veaba/deno-docs/git/commits/4d05eaf539c7ab0b600a4363b52a556ac4fb6317"}, NodeID:"MDM6UmVmMTc1MTQ4ODAxOnJlZnMvaGVhZHMvbWFzdGVy"}
	*/
}

// 获取列表，但无法获判断默认 branch
// func getRefList() {

// 	ctx := context.Background()
// 	ts := oauth2.StaticTokenSource(
// 		&oauth2.Token{AccessToken: "bbbbbbbbbbbbbbbbbb"},
// 	)
// 	tc := oauth2.NewClient(ctx, ts)
// 	client := github.NewClient(tc)

// 	// ref := github.Reference{
// 	// 	Ref:github.String("veaba-bot@123")
// 	// }B
// 	branchOps := &github.BranchListOptions{
// 		Protected:   nil,
// 		ListOptions: github.ListOptions{Page: 1},
// 	}
// 	repos, _, err := client.Repositories.ListBranches(ctx, "veaba", "deno-docs", branchOps)
// 	// client.Repositories.GetCommitSHA1()
// 	fmt.Println("repos=>", &repos)                     // branches [0xc00017d600 0xc0001be0a0 0xc0001be0c0 0xc0001be0e0 0xc0001be120]
// 	fmt.Println("repos getname=>", repos[0].GetName()) // branches [0xc00017d600 0xc0001be0a0 0xc0001be0c0 0xc0001be0e0 0xc0001be120]

// 	for i, v := range repos {
// 		fmt.Println("i=>", i, v.Name, v.GetName(), *v.Name)
// 	}
// 	fmt.Println("err=>", err)
// 	// fmt.Println("name=>", repos.GetLogin())
// 	// fmt.Println("email=>", repos.GetEmail())}
// }

// func createRef() {
// 	ctx := context.Background()
// 	ts := oauth2.StaticTokenSource(
// 		&oauth2.Token{AccessToken: "ccccccccccccccccccc"},
// 	)
// 	tc := oauth2.NewClient(ctx, ts)

// 	fmt.Println("ctx=>", ctx)
// 	fmt.Println("ts=>", ts)
// 	fmt.Println("tc=>", tc)
// 	client := github.NewClient(tc)

// 	// list all repositories for the authenticated user

// 	// contentOps := &github.RepositoryContentGetOptions{
// 	// 	Ref: "main",
// 	// }
// 	ref := github.Reference{
// 		Ref: github.String("refs/heads/veaba-bot@abc"),
// 		Object: &github.GitObject{
// 			SHA: github.String("eeeeeeeeeeeeeeeeee"),
// 		},
// 	}
// 	repos, a, err := client.Git.CreateRef(ctx, "veaba", "deno-docs", &ref)
// 	// client.Repositories.GetCommitSHA1()
// 	fmt.Println("repos=>", repos)
// 	fmt.Println("a=>", a)
// 	fmt.Println("err=>", err)
// 	// fmt.Println("name=>", repos.GetLogin())
// 	// fmt.Println("email=>", repos.GetEmail())
// }
