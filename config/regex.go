package config

import "regexp"

var (
	ListUserRe   = regexp.MustCompile(`^\/users[\/]*$`)
	GetUserRe    = regexp.MustCompile(`^\/users\/(\d+)$`)
	CreateUserRe = regexp.MustCompile(`^\/users[\/]*$`)
)
