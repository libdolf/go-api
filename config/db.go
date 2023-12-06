package config

import (
	"sync"

	"github.com/libdolf/go-api/schemas"
)

type Datastore struct {
	m map[string]schemas.User
	*sync.RWMutex
}
