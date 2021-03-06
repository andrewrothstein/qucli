package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpPut(url string, apitoken string, body []byte) (string, error) {
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	if apitoken != "" {
		req.Header.Set("Authorization", "Bearer "+apitoken)
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP error!\nURL: %s\nstatus code: %d\nbody:\n%s\n", url, resp.StatusCode, string(b))
	}

	return string(b), nil
}
