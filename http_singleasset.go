package main

import (
	"bytes"
	"fmt"
	"io"
	"mime"
	"net/http"
	"path"
)

type singleAssetHandler struct {
	path string
}

func NewSingleAssetHandler(path string) *singleAssetHandler {
	fmt.Printf("NewSingleAssetHandler %s\n", path)
	return &singleAssetHandler{path}
}

func (s *singleAssetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := Asset(s.path)
	if err != nil {
		http.NotFound(w, r)
		fmt.Fprintf(w, "Not found %v", s.path)
		// Asset was not found.
	}
	fmt.Printf("ui request %s\n", s.path)
	w.Header().Set("Content-Type", mime.TypeByExtension(path.Ext(s.path)))
	io.Copy(w, bytes.NewBuffer(data))
}
