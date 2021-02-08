package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "ddddddddddd"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user

	contentOps := &github.RepositoryContentGetOptions{
		Ref: "main",
	}
	repos, a, _, err := client.Repositories.GetContents(ctx, "veaba", "vuepress-plugin-editable", "docs/README.md", contentOps)
	// client.Repositories.GetCommitSHA1()
	// fmt.Println("repos=>", repos)
	fmt.Println("a=>", a)
	fmt.Println("err=>", err)
	fmt.Println("sha=>", repos.SHA)
	fmt.Println("sha1=>", repos.GetSHA())
	content, _ := repos.GetContent()
	fmt.Println("content=>", content)
	var line = 7
	var splitContent = strings.Split(content, "\n")

	splitContent[line] = "哇哈哈哈" + "\n"
	// TODO: line 大于的情况

	ret := ""
	for _, item := range splitContent {
		ret += item
	}
	fmt.Println("new content=>", ret)
	var newStr = strings.Join(splitContent, "\n")
	fmt.Println("newStr content=>", newStr)

}
