package gitlab

import (
	"github.com/xanzy/go-gitlab"
)

type Config struct {
	Token string
	Ref   string
	Pid   int
	Url   string
}

type gitlabInfo struct {
	config Config
	git    *gitlab.Client
}

func NewGitlab(c Config) *gitlabInfo {
	git, _ := gitlab.NewClient(c.Token, gitlab.WithBaseURL(c.Url))
	return &gitlabInfo{
		config: c,
		git:    git,
	}
}

func (g *gitlabInfo) Fetch(fileName string) ([]byte, error) {
	data, _, err := g.git.RepositoryFiles.GetRawFile(g.config.Pid, fileName, &gitlab.GetRawFileOptions{Ref: &g.config.Ref})
	if err != nil {
		return nil, err
	}
	return data, nil
}
