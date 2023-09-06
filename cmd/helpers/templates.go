package helpers

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/accmeboot/issueshift/web"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

type Cache map[string]*template.Template

type RenderDTO struct {
	Writer   http.ResponseWriter
	Template string
	Name     string
	Status   int
	Data     any
}

func formatTime(t time.Time, layout string) string {
	return t.UTC().Format(layout)
}

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

	components, err := fs.Glob(web.Files, "components/*.gohtml")
	if err != nil {
		return nil, err
	}

	for _, component := range components {
		name := filepath.Base(component)

		ts, err := template.New(name).Funcs(funcMap).ParseFS(web.Files, components...)

		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return &cache, nil
}

func (c Cache) Render(DTO RenderDTO) {
	if DTO.Writer == nil {
		panic(errors.New("templates: no response writer"))
	}

	if DTO.Template == "" {
		panic(errors.New("templates: no template name"))
	}

	ts, ok := c[DTO.Template]
	if !ok {
		c.ServerError(DTO.Writer, fmt.Errorf("template %s: does not exist", DTO.Template))
		return
	}

	buffer := new(bytes.Buffer)

	name := "base"
	if DTO.Name != "" {
		name = DTO.Name
	}

	if err := ts.ExecuteTemplate(buffer, name, DTO.Data); err != nil {
		c.ServerError(DTO.Writer, fmt.Errorf("failed to execute template: %s; %s", DTO.Template, err.Error()))
		return
	}

	status := http.StatusOK
	if DTO.Status != 0 {
		status = DTO.Status
	}

	DTO.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	DTO.Writer.WriteHeader(status)
	if _, err := buffer.WriteTo(DTO.Writer); err != nil {
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
