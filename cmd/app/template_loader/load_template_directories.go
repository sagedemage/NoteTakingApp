package template_loader

import (
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
)

func LoadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}

	page_not_found_path := filepath.Join(templatesDir, "errors/404.html")

	r.AddFromFiles("404.html", page_not_found_path)

	unathorized_page_path := filepath.Join(templatesDir, "errors/401.html")

	r.AddFromFiles("401.html", unathorized_page_path)

	return r
}

