package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
	"path/filepath"
)

func main() {
	r := gin.Default()

	r.HTMLRender = loadTemplates("./resources/templates")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.tmpl", nil)
	})

	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	if err != nil {
		log.Fatal(err)
	}
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	templatesSubDirs, err := filepath.Glob(templatesDir + "/*")
	if err != nil {
		log.Fatal(err)
	}

	log.Print("templateSubDir", templatesSubDirs)
	for _, templateSubDir := range templatesSubDirs {
		subdirBase := filepath.Base(templateSubDir)

		if subdirBase == "layouts" {
			continue
		}

		layouts, err := filepath.Glob(path.Join(templatesDir, "layouts", subdirBase+"*.tmpl"))
		log.Print("Layout glob", path.Join(templatesDir, "layouts", subdirBase+"*.tmpl"))
		if err != nil {
			log.Fatal(err)
		}

		pages, err := filepath.Glob(path.Join(templateSubDir, "*.tmpl"))
		if err != nil {
			log.Fatal(err)
		}

		log.Print("layouts", layouts)
		log.Print("pages", templateSubDir, pages)

		// Generate our templates map from each subdirectory
		for _, page := range pages {
			layoutCopy := make([]string, len(layouts))
			copy(layoutCopy, layouts)
			files := append(layoutCopy, page)
			log.Print("Adding:", filepath.Base(page), files)
			r.AddFromFiles(filepath.Base(page), files...)
		}
	}

	return r
}
