package main

import (
	"context"
	"errors"
	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
	"io/ioutil"
	"strings"
	"time"
)

const (
	sourceOwner  = "diov"
	sourceRepo   = "wiki"
	commitBranch = "main"
)

var (
	ctx           = context.Background()
	authorName    = "diov-bot"
	authorEmail   = "bot@mail.dio.wtf"
	commitMessage = "Upgrade gadio rss feed"
)

func pushFeedFile(path, token string) error {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	ref, err := getRef(client)
	if err != nil {
		return err
	}
	if ref == nil {
		return err
	}
	tree, err := getTree(client, ref, path)
	if err != nil {
		return err
	}

	if err := pushCommit(client, ref, tree); err != nil {
		return err
	}
	return nil
}

func getRef(client *github.Client) (ref *github.Reference, err error) {
	ref, _, err = client.Git.GetRef(ctx, sourceOwner, sourceRepo, "refs/heads/"+commitBranch)
	return ref, err
}

func getTree(client *github.Client, ref *github.Reference, path string) (tree *github.Tree, err error) {
	// Create a tree with what to commit.
	var entries []*github.TreeEntry

	file, content, err := getFileContent(path)
	if err != nil {
		return nil, err
	}
	entries = append(entries, &github.TreeEntry{Path: github.String(file), Type: github.String("blob"), Content: github.String(string(content)), Mode: github.String("100644")})
	tree, _, err = client.Git.CreateTree(ctx, sourceOwner, sourceRepo, *ref.Object.SHA, entries)
	return tree, err
}

func getFileContent(fileArg string) (targetName string, b []byte, err error) {
	var localFile string
	files := strings.Split(fileArg, ":")
	switch {
	case len(files) < 1:
		return "", nil, errors.New("empty `-files` parameter")
	case len(files) == 1:
		localFile = files[0]
		targetName = files[0]
	default:
		localFile = files[0]
		targetName = files[1]
	}

	b, err = ioutil.ReadFile(localFile)
	return targetName, b, err
}

func pushCommit(client *github.Client, ref *github.Reference, tree *github.Tree) (err error) {
	// Get the parent commit to attach the commit to.
	parent, _, err := client.Repositories.GetCommit(ctx, sourceOwner, sourceRepo, *ref.Object.SHA, nil)
	if err != nil {
		return err
	}
	// This is not always populated, but is needed.
	parent.Commit.SHA = parent.SHA

	// Create the commit using the tree.
	date := time.Now()
	author := &github.CommitAuthor{Date: &date, Name: &authorName, Email: &authorEmail}
	commit := &github.Commit{Author: author, Message: &commitMessage, Tree: tree, Parents: []*github.Commit{parent.Commit}}
	newCommit, _, err := client.Git.CreateCommit(ctx, sourceOwner, sourceRepo, commit)
	if err != nil {
		return err
	}

	// Attach the commit to the master branch.
	ref.Object.SHA = newCommit.SHA
	_, _, err = client.Git.UpdateRef(ctx, sourceOwner, sourceRepo, ref, false)
	return err
}
