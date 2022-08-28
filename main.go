package main

import (
	"flag"
	"fmt"
	"gophercises/exercise03/cli"
	"gophercises/exercise03/handler"
	"gophercises/exercise03/story"
	"html/template"
	"net/http"
	"path/filepath"
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
