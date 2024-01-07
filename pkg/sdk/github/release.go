package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Masterminds/semver/v3"

	"github.com/voidint/go-update"
)

// Release version
type Release struct {
	TagName string  `json:"tag_name"`
	Assets  []Asset `json:"assets"`
}

// Asset static resources
type Asset struct {
	Name               string `json:"name"`
	ContentType        string `json:"content_type"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

func (a Asset) IsCompressedFile() bool {
	return a.ContentType == "application/zip" || a.ContentType == "application/x-gzip"
}

type ReleaseUpdater struct {
}

func NewReleaseUpdater() *ReleaseUpdater {
	return new(ReleaseUpdater)
}

func (up ReleaseUpdater) CheckForUpdates(current *semver.Version, owner, repo string) (rel *Release, yes bool, err error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, false, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, false, err
	}
	defer resp.Body.Close()

	if !IsSuccess(resp.StatusCode) {
		return nil, false, NewURLUnreachableError(url, fmt.Errorf("%d", resp.StatusCode))
	}

	var latest Release
	if err = json.NewDecoder(resp.Body).Decode(&latest); err != nil {
		return nil, false, err
	}

	latestVersion, err := semver.NewVersion(latest.TagName)
	if err != nil {
		return nil, false, err
	}
	if latestVersion.GreaterThan(current) {
		return &latest, true, nil
	}
	return nil, false, nil
}

var ErrAssetNotFound = errors.New("asset not found")

func (up ReleaseUpdater) Apply(rel *Release,
	findAsset func([]Asset) (idx int),
	proxy string,
) error {
	idx := findAsset(rel.Assets)
	if idx < 0 {
		return ErrAssetNotFound
	}
	tmpDir, err := os.MkdirTemp("", strconv.FormatInt(time.Now().UnixNano(), 10))
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)

	url := rel.Assets[idx].BrowserDownloadURL
	srcFilename := filepath.Join(tmpDir, filepath.Base(url))
	dstFilename := srcFilename
	if _, err = Download(url, proxy, srcFilename, os.O_WRONLY|os.O_CREATE, 0644, true); err != nil {
		return err
	}
	dstFile, err := os.Open(dstFilename)
	if err != nil {
		return nil
	}
	defer dstFile.Close()
	return update.Apply(dstFile, update.Options{})
}
