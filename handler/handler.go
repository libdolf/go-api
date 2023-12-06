package handler

import (
	"net/http"
	"regexp"

	"github.com/libdolf/go-api/api_err"
	"github.com/libdolf/go-api/config"
)

var (
	ListUserRe   = regexp.MustCompile(`^\/users[\/]*$`)
	GetUserRe    = regexp.MustCompile(`^\/users\/(\d+)$`)
	CreateUserRe = regexp.MustCompile(`^\/users[\/]*$`)
)

type UserHandler struct {
	store *config.Datastore
}

func NewUserHandler(s *config.Datastore) *UserHandler {
	return &UserHandler{store: s}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "aplication/json")
	switch {
	case r.Method == http.MethodGet && ListUserRe.MatchString(r.URL.Path):
		h.List(w, r)
		return
	case r.Method == http.MethodGet && GetUserRe.MatchString(r.URL.Path):
		h.Get(w, r)
		return
	case r.Method == http.MethodPost && CreateUserRe.MatchString(r.URL.Path):
		h.Create(w, r)
		return
	default:
		api_err.NotFound(w, r)
		return
	}
}
