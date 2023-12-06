package handler

import (
	"encoding/json"
	"net/http"

	"github.com/libdolf/go-api/api_err"
	"github.com/libdolf/go-api/schemas"
)

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	h.store.RLock()
	users := make([]schemas.User, 0, len(h.store.M))
	for _, v := range h.store.M {
		users = append(users, v)
	}
	h.store.RUnlock()
	jsonBytes, err := json.Marshal(users)
	if err != nil {
		api_err.InternalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
