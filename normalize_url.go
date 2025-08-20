package main

import (
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	host := parsedURL.Host
	if host == "" {
		parsedURL, err = url.Parse("https://" + rawURL)
		if err != nil {
			return "", err
		}
		host = parsedURL.Host
	}
	path := strings.TrimSuffix(parsedURL.EscapedPath(), "/")
	return host + path, nil
}
