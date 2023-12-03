package static

import (
	"github.com/kenowi-dev/hsp-auto-login/pkg/api/files"
	"net/http"
	"strings"
)

type Handler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type handler struct {
}

func New() Handler {
	return &handler{}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {

	// remove leading "/" so it maps onto the static folder
	file, err := files.Static.ReadFile(strings.TrimPrefix(r.URL.Path, "/"))
	if err != nil {
		panic(err)
	}
	_, err = w.Write(file)
	if err != nil {
		panic(err)
	}
}
