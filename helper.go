package main

import (
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func get_fname(resp *http.Response, raw string) string {
	if cd := resp.Header.Get("Content-Disposition"); cd != "" {
		_, params, err := mime.ParseMediaType(cd)
		if err == nil {
			if name, ok := params["filename*"]; ok {
				if decoded, err := url.QueryUnescape(strings.TrimPrefix(name, "UTF-8''")); err == nil {
					return decoded
				}
			}
		}
	}
	u, err := url.Parse(raw)
	if err == nil {
		name := path.Base(u.Path)
		if name != "/" && name != "." {
			return name
		}
	}

	return ""
}

func download(url string, dwnld_path string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Bad Status Code %d", res.StatusCode)
	}

	fname := get_fname(res, url)
	if fname == "" {
		return errors.New("Unable to parse fname")
	}
	fpath := dwnld_path
	if err = os.MkdirAll(fpath, 0755); err != nil {
		return fmt.Errorf("Failed to create download directory: %w", err)
	}

	fpath = filepath.Join(fpath, fname)
	fmt.Println(fpath)
	out, err := os.Create(fpath)
	if err != nil {
		return errors.New("Failed to create file" + fpath)
	}
	defer out.Close()
	_, err = io.Copy(out, res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func parse_args(args []string) map[string]string {
	ret := make(map[string]string)
	prev := ""
	for i, val := range args {
		if i&1 != 0 {
			ret[prev] = val
		} else {
			prev = val
			ret[val] = ""
		}
	}
	return ret
}

func usage(argv string) {
	fmt.Println("Usage: ", argv, " [URL DEST]...")
}
