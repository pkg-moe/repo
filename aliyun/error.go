package aliyun

import repo_error "pkg.moe/pkg/repo"

var (
	// ErrInvalidInteger defines invalid integer error
	ErrInvalidInteger = repo_error.ErrInvalidInteger
	// ErrInvalidJSON defines invalid json error
	ErrInvalidJSON = repo_error.ErrInvalidJSON
	// ErrInvalidType defines invalid type error
	ErrInvalidType = repo_error.ErrInvalidType
	//ErrNoRowsFound describes the error
	ErrNoRowsFound = repo_error.ErrNoRowsFound
	// ErrInvalidWeChatConf describes the error
	ErrInvalidConf = repo_error.ErrInvalidConf
)
