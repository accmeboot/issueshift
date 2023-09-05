package helpers

import (
	"bytes"
	"github.com/accmeboot/issueshift/web"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

type Cache map[string]*template.Template

func formatTime(t time.Time, layout string) string {
	return t.UTC().Format(layout)
}

//"02 Jan 2006 at 15:04"

var funcMap = template.FuncMap{
	"formatTime": formatTime,
}

func NewCache() (*Cache, error) {
	cache := Cache{}

	pages, err := fs.Glob(web.Files, "pages/*.gohtml")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		patterns := []string{
			"components/*.gohtml",
			page,
		}

		ts, err := template.New(name).Funcs(funcMap).ParseFS(web.Files, patterns...)

		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return &cache, nil
}

func (c Cache) Render(w http.ResponseWriter, status int, page string, fragment *string, data any) {
	ts, ok := c[page]
	if !ok {
		log.Printf("template %s: does not exist", page)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	// First executing a template into a buffer, in order to prevent corrupted templates to a client
	buf := new(bytes.Buffer)
	tmlFragment := "base"

	if fragment != nil {
		tmlFragment = *fragment
	}

	err := ts.ExecuteTemplate(buf, tmlFragment, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	if _, err = buf.WriteTo(w); err != nil {
		log.Println(err)
	}
}

func (c Cache) ServerError(w http.ResponseWriter, err error) {
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("HX-Redirect", "/error")
	w.WriteHeader(http.StatusInternalServerError)
}
