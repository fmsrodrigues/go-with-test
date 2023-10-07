package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

var (
	//go:embed templates/*
	postTemplates embed.FS
)

type Post struct {
	Title, SanitisedTitle, Body, Description string
	Tags                                     []string
}

type PostRenderer struct {
	templ *template.Template
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", p)
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}
