package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/sadex11/gopher03adventure/cli"
	"github.com/sadex11/gopher03adventure/handler"
	"github.com/sadex11/gopher03adventure/story"
)

func main() {
	runCli := flag.Bool("cli", true, "True - run in CLI, False - run as HTTP server")

	flag.Parse()
	storyPath, err := filepath.Abs("gopher.json")

	if err != nil {
		panic(err)
	}

	storyMap := story.CreateStories(storyPath)

	if *runCli {
		cliHandler := cli.NewCliHandler(storyMap)
		cliHandler.RunCli()
	} else {
		tmpl := template.Must(template.ParseFiles("template.html"))
		httpHandler := handler.NewStoryHandler(storyMap, tmpl)

		fmt.Println("Run HTTP server on port 8080")
		http.ListenAndServe(":8080", httpHandler)
	}

}
