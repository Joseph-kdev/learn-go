package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"strings"

	"github.com/yuin/goldmark"
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
	HTMLBody                 template.HTML
}

type PostRenderer struct {
	templ *template.Template
}

type PostViewModel struct {
	Title, SanitisedTitle, Description, Body string
	Tags                                     []string
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS

	md = goldmark.New(
		goldmark.WithExtensions(),
	)
)

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	p = PreparePost(p)
	return r.templ.ExecuteTemplate(w, "blog.gohtml", p)
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}

func markdownToHTML(markdown string) template.HTML {
	var buf strings.Builder
	if err := md.Convert([]byte(markdown), &buf); err != nil {
		return template.HTML(markdown)
	}
	return template.HTML(buf.String())
}

func PreparePost(p Post) Post {
	p.HTMLBody = markdownToHTML(p.Body)
	return p
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}
