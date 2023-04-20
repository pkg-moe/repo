package repo_model

type RepoSourceModel interface {
	Fetch(filename string) ([]byte, error)
}
