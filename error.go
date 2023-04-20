package repo_error

import "errors"

var (
	// ErrInvalidInteger defines invalid integer error
	ErrInvalidInteger = errors.New("invalid integer")
	// ErrInvalidJSON defines invalid json error
	ErrInvalidJSON = errors.New("invalid json")
	// ErrInvalidType defines invalid type error
	ErrInvalidType = errors.New("invalid type")
	//ErrNoRowsFound describes the error
	ErrNoRowsFound = errors.New("No Rows Found")
	// ErrInvalidWeChatConf describes the error
	ErrInvalidConf = errors.New("Invalid Conf")
	// ErrInvalidGitFile describes the error of invalid git file
	ErrInvalidGitFile = errors.New("invalid git file")
)
