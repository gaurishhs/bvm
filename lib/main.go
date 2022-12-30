package lib

import (
	"context"
	"runtime"

	"github.com/google/go-github/github"
)

func GetAllReleases() ([]*github.RepositoryRelease, error) {
	releases, _, err := client.Repositories.ListReleases(context.Background(), "oven-sh", "bun", nil)
	if err != nil {
		return nil, err
	}
	return releases, nil
}

func GetLatestRelease() (*github.RepositoryRelease, error) {
	release, _, err := client.Repositories.GetLatestRelease(context.Background(), "oven-sh", "bun")
	if err != nil {
		return nil, err
	}
	return release, nil
}

func GetPlatform() string {
	if runtime.GOOS == "windows" {
		return "win"
	} else if runtime.GOOS == "linux" {
		return "linux"
	} else if runtime.GOOS == "darwin" {
		return "darwin"
	} else {
		return "other"
	}
}

func GetArch() string {
	if runtime.GOARCH == "arm64" {
		return "aarch64"
	} else if runtime.GOARCH == "amd64" {
		return "x64"
	} else {
		return "other"
	}
}

func InitalizeClient() {
	client = github.NewClient(nil)
}

var (
	client *github.Client
)
