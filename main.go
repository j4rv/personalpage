package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"log"
	"path"
	"path/filepath"
)

func main() {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	r.HTMLRender = loadTemplates("./resources/templates")

	r.Static("/resources/static", "./resources/static")

	r.GET("/", indexHandler)

	r.NoRoute(notFoundHandler)

	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	if err != nil {
		log.Fatal(err)
	}
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	subDirs, err := filepath.Glob(templatesDir + "/*")
	if err != nil {
		log.Fatal(err)
	}

	for _, subDir := range subDirs {
		subDirBase := filepath.Base(subDir)

		if subDirBase == "layouts" {
			continue
		}

		// Get the layouts for the current directory
		// Layouts must begin with the Name of the directory and end in .tmpl
		layouts, err := filepath.Glob(path.Join(templatesDir, "layouts", subDirBase+"*.tmpl"))
		if err != nil {
			log.Fatal(err)
		}

		// Get all .tmpl pages from the current directory
		pages, err := filepath.Glob(path.Join(subDir, "*.tmpl"))
		if err != nil {
			log.Fatal(err)
		}

		// Generate our templates map for the current directory
		for _, page := range pages {
			layoutCopy := make([]string, len(layouts))
			copy(layoutCopy, layouts)
			files := append(layoutCopy, page)
			r.AddFromFiles(filepath.Base(page), files...)
		}
	}

	return r
}
