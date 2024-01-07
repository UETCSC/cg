package github

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"time"

	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
	"golang.org/x/net/proxy"
)

// Download the resource and save as
func Download(srcURL string, proxy_str string, filename string, flag int, perm fs.FileMode, withProgress bool) (size int64, err error) {
	httpClient := &http.Client{}
	if proxy_str != "" {
		fmt.Println("使用代理 socks5:// " + proxy_str)
		dialer, err := proxy.SOCKS5("tcp", proxy_str, nil, proxy.Direct)
		if err != nil {
			fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
			os.Exit(1)
		}
		// Setup an http client
		httpTransport := &http.Transport{}

		httpTransport.Dial = dialer.Dial
		httpClient = &http.Client{Transport: httpTransport}
	}

	//Send a web request
	resp, err := httpClient.Get(srcURL)
	if err != nil {
		return 0, NewDownloadError(srcURL, err)
	}
	defer resp.Body.Close()

	if !IsSuccess(resp.StatusCode) {
		return 0, NewURLUnreachableError(srcURL, fmt.Errorf("%d", resp.StatusCode))
	}

	f, err := os.OpenFile(filename, flag, perm)
	if err != nil {
		return 0, NewDownloadError(srcURL, err)
	}
	defer f.Close()

	var dst io.Writer
	if withProgress {
		bar := progressbar.NewOptions64(
			resp.ContentLength,
			progressbar.OptionEnableColorCodes(true),
			progressbar.OptionSetTheme(progressbar.Theme{
				Saucer:        "=",
				SaucerHead:    ">",
				SaucerPadding: " ",
				BarStart:      "[",
				BarEnd:        "]",
			}),
			progressbar.OptionSetWidth(15),
			progressbar.OptionSetDescription("Downloading"),
			progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
			progressbar.OptionShowBytes(true),
			progressbar.OptionThrottle(65*time.Millisecond),
			progressbar.OptionShowCount(),
			progressbar.OptionOnCompletion(func() {
				_, _ = fmt.Fprint(ansi.NewAnsiStdout(), "\n")
			}),
		)
		_ = bar.RenderBlank()
		dst = io.MultiWriter(f, bar)

	} else {
		dst = f
	}
	return io.Copy(dst, resp.Body)
}

// IsSuccess returns the request to http for success
func IsSuccess(statusCode int) bool {
	return statusCode >= http.StatusOK && statusCode < http.StatusMultipleChoices
}
