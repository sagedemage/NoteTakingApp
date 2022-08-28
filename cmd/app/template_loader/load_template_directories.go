package template_loader

import (
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
)

func generate_templates(r multitemplate.Renderer, templatesDir string, layoutFile string, pagesDir string) multitemplate.Renderer {
	layouts, err := filepath.Glob(templatesDir + layoutFile)
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + pagesDir)
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
	return r
}

func LoadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	r = generate_templates(r, templatesDir, "/layouts/base.tmpl", "/includes/*.tmpl")

	r = generate_templates(r, templatesDir, "/layouts/error-base.html", "/errors/*.html")

	return r
}

