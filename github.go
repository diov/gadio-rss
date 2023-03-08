package main

import (
	"context"
	"errors"
	"io/ioutil"
	"strings"
	"time"

	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

const (
	sourceOwner  = "diov"
	sourceRepo   = "gadio-rss"
	targetRepo   = "wiki"
	commitBranch = "main"
)

var (
	ctx           = context.Background()
	authorName    = "diov-bot"
	authorEmail   = "bot@mail.dio.wtf"
	commitMessage = "Upgrade gadio rss feed"
	tempFilePath  = "target/artifact.zip"
	gitMgr        *gitManager
)

type gitManager struct {
	client *github.Client
}

func setupGitManager(token string) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	gitMgr = &gitManager{client: client}
}

func (m *gitManager) getPreviousArtifact() error {
	if nil == gitMgr {
		return nil
	}

	opts := &github.ListOptions{Page: 1}
	artifacts, _, err := m.client.Actions.ListArtifacts(ctx, sourceOwner, sourceRepo, opts)
	if nil != err {
		return err
	}
	if artifacts.GetTotalCount() <= 0 {
		return nil
	}
	artifact := artifacts.Artifacts[0]
	url, _, err := gitMgr.client.Actions.DownloadArtifact(ctx, sourceOwner, sourceRepo, artifact.GetID(), true)
	if nil != err {
		return err
	}
	err = downloadFile(url.String(), tempFilePath)
	if nil != err {
		return err
	}
	err = unzipFile(tempFilePath)
	if nil != err {
		return err
	}
	return err
}

func (m *gitManager) pushFeedFile(path string) error {
	if nil == gitMgr {
		return nil
	}

	ref, err := m.getRef()
	if err != nil {
		return err
	}
	if ref == nil {
		return err
	}
	tree, err := m.getTree(ref, path)
	if err != nil {
		return err
	}

	if err := m.pushCommit(ref, tree); err != nil {
		return err
	}
	return nil
}

func (m *gitManager) getRef() (ref *github.Reference, err error) {
	ref, _, err = m.client.Git.GetRef(ctx, sourceOwner, targetRepo, "refs/heads/"+commitBranch)
	return ref, err
}

func (m *gitManager) getTree(ref *github.Reference, path string) (tree *github.Tree, err error) {
	// Create a tree with what to commit.
	var entries []*github.TreeEntry

	file, content, err := m.getFileContent(path)
	if err != nil {
		return nil, err
	}
	entries = append(entries, &github.TreeEntry{Path: github.String(file), Type: github.String("blob"), Content: github.String(string(content)), Mode: github.String("100644")})
	tree, _, err = m.client.Git.CreateTree(ctx, sourceOwner, targetRepo, *ref.Object.SHA, entries)
	return tree, err
}

func (m *gitManager) getFileContent(fileArg string) (targetName string, b []byte, err error) {
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

func (m *gitManager) pushCommit(ref *github.Reference, tree *github.Tree) (err error) {
	// Get the parent commit to attach the commit to.
	parent, _, err := m.client.Repositories.GetCommit(ctx, sourceOwner, targetRepo, *ref.Object.SHA, nil)
	if err != nil {
		return err
	}
	// This is not always populated, but is needed.
	parent.Commit.SHA = parent.SHA

	// Create the commit using the tree.
	date := time.Now()
	author := &github.CommitAuthor{Date: &date, Name: &authorName, Email: &authorEmail}
	commit := &github.Commit{Author: author, Message: &commitMessage, Tree: tree, Parents: []*github.Commit{parent.Commit}}
	newCommit, _, err := m.client.Git.CreateCommit(ctx, sourceOwner, targetRepo, commit)
	if err != nil {
		return err
	}

	// Attach the commit to the master branch.
	ref.Object.SHA = newCommit.SHA
	_, _, err = m.client.Git.UpdateRef(ctx, sourceOwner, targetRepo, ref, false)
	return err
}
