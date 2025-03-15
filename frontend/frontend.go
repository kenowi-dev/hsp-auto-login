package frontend

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
)

// /go:generate bun install --cwd ./frontend
// /go:generate bun run --cwd ./frontend build
//
//go:generate deno install
//go:generate deno task build
//go:embed all:build
var files embed.FS
var svelteFiles = must(fs.Sub(files, "build"))

var defaultHandler = StaticHandler("/")

func StaticIndexHandlerFunc(w http.ResponseWriter, r *http.Request) {
	defaultHandler.ServeHTTP(w, r)
}

func StaticHandler(path string) http.HandlerFunc {
	prefix := path
	if path == "" {
		prefix = "/"
	}

	return func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Path
		// escape url path
		tmpPath, err := url.PathUnescape(p)
		if err != nil {
			http.Error(w, fmt.Errorf("failed to unescape path variable: %w", err).Error(), http.StatusInternalServerError)
			return
		}
		p = tmpPath

		// fs.FS.Open() already assumes that file names are relative to FS root path and considers name with prefix `/` as invalid
		name := filepath.ToSlash(filepath.Clean(strings.TrimPrefix(p, prefix)))

		if name == "" || name == "." {
			http.ServeFileFS(w, r, svelteFiles, "index.html")
			return
		}

		switch filepath.Ext(name) {
		case ".js", ".css", ".png":
			http.ServeFileFS(w, r, svelteFiles, name)
		case ".html":
			// disallow routes with explicit .html extension
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		case "":
			http.ServeFileFS(w, r, svelteFiles, name+".html")
		default:
			http.ServeFileFS(w, r, svelteFiles, name)
		}
	}
}

func must[T any](obj T, err error) T {
	if err != nil {
		panic(err)
	}
	return obj
}
