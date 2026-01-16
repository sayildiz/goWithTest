package blogrenderer

import (
	"embed"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"html/template"
	"io"
	"strings"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

var (
	//this is a directive used to import an embed to init a variable of type string, []byte, or FS
	//with the contents of files read from the package directory or subdirectories at compile time
	//go:embed "templates/*"
	postTemplate embed.FS
)

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplate, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", newPostVM(p, r))
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}

type postViewModel struct {
	Post
	HTMLBody template.HTML
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}

func newPostVM(p Post, r *PostRenderer) postViewModel {
	vm := postViewModel{Post: p}
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	parser := parser.NewWithExtensions(extensions)
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(p.Body), parser, nil))
	return vm
}
