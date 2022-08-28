package handler

import (
	"fmt"
	"github.com/sadex11/gopher03adventure/story"
	"html/template"
	"net/http"
	"strings"
)

type StoryHandler struct {
	StoryMap      story.StoryConfig
	StoryTemplate *template.Template
}

func NewStoryHandler(storyMap story.StoryConfig, tmpl *template.Template) StoryHandler {
	return StoryHandler{storyMap, tmpl}
}

func getStoryPath(requestPath string) string {
	parsed := strings.TrimSpace(requestPath)
	parsed = strings.TrimPrefix(parsed, "/")

	// set base path of the story
	if parsed == "" {
		parsed = "intro"
	}

	return parsed
}

func (h StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	storyPath := getStoryPath(r.URL.Path)

	if chapter, ok := h.StoryMap[storyPath]; ok {
		err := h.StoryTemplate.Execute(w, chapter)

		if err != nil {
			fmt.Println("Path:", storyPath, " - error:", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Story chapter not found!"))
	}
}
