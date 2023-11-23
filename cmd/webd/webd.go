package main

import (
	"github.com/Masterminds/sprig"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/spf13/viper"
	"go-web-base/cmd/webd/routes"
	"go-web-base/internal/config"
	"go-web-base/internal/types"
	"go-web-base/internal/webpack"
	"go-web-base/internal/wire"
	"html/template"
	"io/fs"
	"log"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	err := config.Setup()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	webpack.Init(viper.GetString("env") == "development")

	router := gin.Default()
	router.HTMLRender = loadTemplates("templates")
	router.Use(func(ctx *gin.Context) {
		if strings.HasPrefix(ctx.Request.URL.Path, "/dist/") {
			ctx.Header("Cache-Control", "public, max-age=15552000")
		}
		ctx.Next()
	})
	router.Static("/dist", "dist")

	todos := []string{
		"Go to the Store",
		"Finish the Website",
		"Go to the Gym",
	}

	todoStore := map[string]*types.Todo{}

	for _, todo := range todos {
		id := uuid.NewString()
		todoStore[id] = &types.Todo{
			ID:      id,
			Created: time.Now().UTC(),
			Updated: time.Now().UTC(),
			Body:    todo,
		}
	}

	r := routes.NewService(routes.Service{
		TodoHandler: wire.InitializeTodoHandler(todoStore),
	})

	r.LoadAppRoutes(router)
	r.LoadPublicRoutes(router)

	log.Fatal(router.Run(":3000"))
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	defaultLayout := templatesDir + "/layouts/default-layout.html"
	appLayout := templatesDir + "/layouts/app-layout.html"

	pagesDir := templatesDir + "/pages"
	fragmentsDir := templatesDir + "/fragments"

	var paths = map[string]string{
		pagesDir + "/public": defaultLayout,
		pagesDir + "/app":    appLayout,
	}

	funcMap := template.FuncMap{
		"Slugify": slug.Make,
		"asset":   webpack.AssetHelper,
	}
	for k, v := range sprig.FuncMap() {
		funcMap[k] = v
	}

	// Process Fragment Templates
	var fragments []string
	err := filepath.WalkDir(fragmentsDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if !d.IsDir() {
			fragments = append(fragments, path)
		}

		return nil
	})
	if err != nil {
		return nil
	}

	for _, fragment := range fragments {
		r.AddFromFilesFuncs(strings.TrimPrefix(filepath.ToSlash(strings.TrimSuffix(fragment, ".html")), templatesDir+"/"), funcMap, fragment)
	}

	// Process & Build Template Pages

	for path, layout := range paths {
		var pages []string
		var scopedFragments []string
		err = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return nil
			}

			if !d.IsDir() {
				pages = append(pages, path)
			}

			if d.IsDir() && d.Name() == "fragments" {
				err = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
					if err != nil {
						return nil
					}

					if !d.IsDir() {
						scopedFragments = append(scopedFragments, path)
					}

					return nil
				})
				if err != nil {
					return nil
				}
			}

			return nil
		})
		if err != nil {
			return nil
		}

		for _, fragment := range scopedFragments {
			tmpls := []string{fragment}
			tmpls = append(tmpls, fragments...)

			r.AddFromFilesFuncs(strings.TrimPrefix(filepath.ToSlash(strings.TrimSuffix(fragment, ".html")), path+"/"), funcMap, tmpls...)
		}

		for _, page := range pages {
			var files []string

			var pageFragments []string
			pageFragments = append(pageFragments, scopedFragments...)
			pageFragments = append(pageFragments, fragments...)

			files = append(files, layout, page)
			files = append(files, layout, page)
			files = append(files, pageFragments...)

			r.AddFromFilesFuncs(strings.TrimPrefix(filepath.ToSlash(strings.TrimSuffix(page, ".html")), pagesDir+"/"), funcMap, files...)
		}
	}

	return r
}
