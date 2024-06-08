package util

import (
	"context"
	"fmt"
	"github.com/google/go-github/v39/github"
)

var githubClient *github.Client

func init() {
	githubClient = github.NewClient(nil)
}

func GetReleaseAssetURL(owner, repo, version, fileName string) (string, error) {
	ctx := context.Background()

	var release *github.RepositoryRelease
	var err error
	if version != "" {
		release, _, err = githubClient.Repositories.GetReleaseByTag(ctx, owner, repo, version)
		if err != nil {
			return "", fmt.Errorf("failed to get release for version %s: %v", version, err)
		}
	} else {
		releases, _, err := githubClient.Repositories.ListReleases(ctx, owner, repo, nil)
		if err != nil {
			return "", fmt.Errorf("failed to list releases: %v", err)
		}
		if len(releases) == 0 {
			return "", fmt.Errorf("no releases found")
		}
		release = releases[0]
	}

	assets, _, err := githubClient.Repositories.ListReleaseAssets(ctx, owner, repo, release.GetID(), nil)
	if err != nil {
		return "", fmt.Errorf("failed to list release assets: %v", err)
	}

	for _, asset := range assets {
		if asset.GetName() == fileName {
			return asset.GetBrowserDownloadURL(), nil
		}
	}

	return "", fmt.Errorf("file '%s' not found in release '%s'", fileName, release.GetTagName())
}

func ListRelease(owner, repo string) ([]*github.RepositoryRelease, error) {
	ctx := context.Background()
	releases, _, err := githubClient.Repositories.ListReleases(ctx, owner, repo, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list releases: %v", err)
	}
	return releases, nil
}
