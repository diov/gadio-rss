package main

import (
	"context"

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
	tempFilePath = "target/artifact.zip"
	gitMgr       *gitManager
)

type gitManager struct {
	client *github.Client
}

func setupGitManager(token string) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)
	client := github.NewClient(tc)
	gitMgr = &gitManager{client: client}
}

func (m *gitManager) getPreviousArtifact() error {
	if nil == m {
		return nil
	}

	opts := &github.ListOptions{Page: 1}
	artifacts, _, err := m.client.Actions.ListArtifacts(context.Background(), sourceOwner, sourceRepo, opts)
	if nil != err {
		return err
	}
	if artifacts.GetTotalCount() <= 0 {
		return nil
	}
	artifact := artifacts.Artifacts[0]
	url, _, err := m.client.Actions.DownloadArtifact(context.Background(), sourceOwner, sourceRepo, artifact.GetID(), true)
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
