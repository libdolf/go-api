package handler

import (
	"encoding/json"
	"net/http"

	"github.com/libdolf/go-api/api_err"
	"github.com/libdolf/go-api/schemas"
)

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var u schemas.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		api_err.InternalServerError(w, r)
		return
	}
	h.store.Lock()
	h.store.M[u.ID] = u
	h.store.Unlock()
	jsonBytes, err := json.Marshal(u)
	if err != nil {
		api_err.InternalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
