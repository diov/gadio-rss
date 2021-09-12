package main

import (
	"archive/zip"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"unicode"
)

func isUrlValid(str string) bool {
	_, err := url.ParseRequestURI(str)
	return nil == err
}

func shouldStop(radios []*Radio) bool {
	if len(radios) < pageSize {
		return true
	}

	for _, radio := range radios {
		data, _ := dbMgr.Find([]byte(radio.ID))
		if len(data) > 0 {
			return true
		}
	}
	return false
}

func writeFile(path string, content []byte) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	_, err = f.Write(content)
	return err
}

func remoteContentLength(url string) int64 {
	client := http.DefaultClient
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", userAgent)
	resp, err := client.Do(req)
	if nil != err {
		return 0
	}
	defer func() { _ = resp.Body.Close() }()
	return resp.ContentLength
}

func cleanString(str string) string {
	clean := strings.Map(func(r rune) rune {
		if unicode.IsGraphic(r) && unicode.IsPrint(r) {
			return r
		}
		return -1
	}, str)
	return clean
}

func fileSize(path string) int64 {
	fi, err := os.Stat(path)
	if err != nil {
		return 0
	}
	// get the size
	return fi.Size()
}

func downloadFile(url, dest string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer func() {
		_ = out.Close()
		_ = resp.Body.Close()
	}()
	_, err = io.Copy(out, resp.Body)
	return err
}

func unzipFile(path string) error {
	archive, err := zip.OpenReader(path)
	if nil != err {
		return err
	}
	defer func() { _ = archive.Close() }()

	for _, file := range archive.File {
		dstFile, err := os.OpenFile(file.Name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if nil != err {
			return err
		}
		fileInArchive, err := file.Open()
		if nil != err {
			return err
		}
		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			return err
		}
		_ = dstFile.Close()
		_ = fileInArchive.Close()
	}
	return nil
}
