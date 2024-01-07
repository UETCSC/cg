package cmdutil

import (
	"cg/pkg/util"
	"fmt"
	"os"
	"runtime"
	"strings"

	"cg/pkg/sdk/github"

	"github.com/Masterminds/semver/v3"
)

const (
	ORG  = "UETCSC"
	REPO = "cg"
)

func Upgrade(proxy string) {
	up := github.NewReleaseUpdater()
	latest, yes, err := up.CheckForUpdates(semver.MustParse(util.Version), ORG, REPO)
	if err != nil {
		os.Exit(1)
	}
	if !yes {
		fmt.Println("cg version is the latest, " + util.Version)
	} else {
		fmt.Printf("cg has a new version %s ", latest.TagName)
		if err = up.Apply(latest, findAsset, proxy); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("cg updated successfully.")
	}
}

func SelfCheck(proxy string) {
	up := github.NewReleaseUpdater()
	latest, _, err := up.CheckForUpdates(semver.MustParse(util.Version), ORG, REPO)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("cg latest version: " + latest.TagName)
	fmt.Println("cg current version: " + util.Version)
}
func findAsset(items []github.Asset) (idx int) {
	suffix := fmt.Sprintf("cg_%s_%s", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("Start downloading %s\n", suffix)
	for i := range items {
		if strings.HasSuffix(items[i].BrowserDownloadURL, suffix) {
			return i
		}
	}
	return -1
}
