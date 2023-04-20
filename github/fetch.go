package github

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type gitFileUrl struct {
	Name string `json:"name"`
	URL  string `json:"git_url"`
}

type gitFileBlob struct {
	Content string `json:"content"`
}

type GithubInfo struct {
	username string
	token    string
	repoURL  string
}

func NewGitHub(username, token, repoURL string) *GithubInfo {
	if ref := os.Getenv("REPO_REF"); ref != "" {
		repoURL = fmt.Sprintf("%s?ref=%s", repoURL, ref)
	}

	return &GithubInfo{
		username: username,
		token:    token,
		repoURL:  repoURL,
	}
}

func (g *GithubInfo) Fetch(filename string) ([]byte, error) {
	fileURL, err := g.fetchGitUrl(filename)
	if err != nil {
		return nil, err
	}

	client := new(http.Client)
	req, err := http.NewRequest(http.MethodGet, fileURL, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(g.username, g.token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if len(body) == 0 {
		return nil, ErrInvalidJSON
	}

	gitFileBlob := new(gitFileBlob)
	if err := json.Unmarshal(body, gitFileBlob); err != nil {
		return nil, err
	}

	data, err := base64.StdEncoding.DecodeString(gitFileBlob.Content)
	if err != nil {
		return nil, err
	}

	if !json.Valid(data) {
		return nil, ErrInvalidJSON
	}
	return data, nil
}

func (g *GithubInfo) fetchGitUrl(filename string) (string, error) {
	client := new(http.Client)
	req, err := http.NewRequest(http.MethodGet, g.repoURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to fetch repo %s: %v", g.repoURL, err)
	}

	req.SetBasicAuth(g.username, g.token)

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch repo %s: %v", g.repoURL, err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read repo %s: %v", g.repoURL, err)
	}
	defer resp.Body.Close()

	if len(body) == 0 {
		return "", fmt.Errorf("repo %s is absent", g.repoURL)
	}

	if !json.Valid(body) {
		return "", fmt.Errorf("repo %s error: %v", g.repoURL, ErrInvalidJSON)
	}

	var gitFileUrls []*gitFileUrl
	if err := json.Unmarshal(body, &gitFileUrls); err != nil {
		return "", fmt.Errorf("failed to json decode %s: %v", string(body), err)
	}

	for _, gitFileUrl := range gitFileUrls {
		if gitFileUrl.Name == filename {
			return gitFileUrl.URL, nil
		}
	}

	return "", fmt.Errorf("file %s error: %v", filename, ErrInvalidGitFile)
}
