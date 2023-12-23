package utils

import (
	"io"
	"net/http"
	"strings"
)

func GetRedirect(redirect string) (string, error) {
	resp, err := http.Get("https://gist.githubusercontent.com/vclemenzi/9dc86c53e12266d9a91dd33ff35fc755/raw/78c8a32897e24a36843682a82216b9f1342ab5c5/redirects.txt")

	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	bodyString := string(body)

	// File Format
	// redirect_name: url

	lines := strings.Split(bodyString, "\n")

	defer resp.Body.Close()

	for _, line := range lines {
		if strings.Contains(line, redirect) {
			return strings.Split(line, ": ")[1], nil
		}
	}

	return "", nil
}
