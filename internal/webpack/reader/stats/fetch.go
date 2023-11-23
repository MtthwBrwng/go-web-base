package stats

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

func getURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("go-webpack: Unexpected status code: %d. Expecting %d", resp.StatusCode, http.StatusOK)
	}

	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}

func devManifest(host, webPath string) ([]byte, error) {
	manifestURL := fmt.Sprint("http://", host, "/", webPath, "/manifest.json")
	body, err := getURL(manifestURL)
	if err != nil {
		return []byte{}, errors.Wrap(err, fmt.Sprintf("go-webpack: Error when loading manifest from url %s", manifestURL))

	}
	return []byte(body), nil
}

func prodManifest(fsPath string) ([]byte, error) {
	manifestPath := "./" + fsPath + "/manifest.json"
	body, err := ioutil.ReadFile(manifestPath)
	if err != nil {
		return []byte{}, errors.Wrap(err, fmt.Sprintf("go-webpack: Error when loading manifest from file %s", manifestPath))
	}
	return body, nil
}
