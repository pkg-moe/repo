package github

import (
	"pkg.moe/pkg/repo"
)

var (
	// ErrInvalidInteger defines invalid integer error
	ErrInvalidInteger = repo_error.ErrInvalidInteger
	// ErrInvalidJSON defines invalid json error
	ErrInvalidJSON = repo_error.ErrInvalidJSON
	// ErrInvalidType defines invalid type error
	ErrInvalidType = repo_error.ErrInvalidType
	//ErrNoRowsFound describes the error
	ErrNoRowsFound = repo_error.ErrNoRowsFound
	// ErrInvalidWeChatConf describes the error of lack of APPID / SECRET
	ErrInvalidConf = repo_error.ErrInvalidConf
	// ErrInvalidGitFile describes the error of invalid git file
	ErrInvalidGitFile = repo_error.ErrInvalidGitFile
)
