package cli

import (
	"fmt"

	"github.com/sadex11/gopher03adventure/story"
)

const (
	Reset = "\033[0m"
	Green = "\033[32m"
	Bold  = "\033[1m"
	Clean = "\033[H\033[2J"
)

type CliHandler struct {
	StoryMap story.StoryConfig
}

func NewCliHandler(storyMap story.StoryConfig) CliHandler {
	return CliHandler{storyMap}
}

func (c CliHandler) getChapter(title string) story.Chapter {
	chapter, ok := c.StoryMap[title]

	if !ok {
		panic(fmt.Sprintf("Chapter %s not found!", title))
	}

	return chapter
}

func (c CliHandler) RunCli() {
	currChapter := c.getChapter("intro")

	for {
		fmt.Println(Clean)
		fmt.Println(Green, currChapter.Title, Reset)

		for _, storyLine := range currChapter.Story {
			fmt.Println(storyLine)
		}

		if len(currChapter.Options) == 0 {
			return
		}

		fmt.Println("\nNext steps: ")

		for i, option := range currChapter.Options {
			fmt.Println(Bold, "[", i+1, "] ", Reset, option.Text)
		}

		fmt.Println("\nPress step number to continue")

		for {
			var i int
			_, err := fmt.Scan(&i)

			if err != nil || i > len(currChapter.Options) || i < 1 {
				fmt.Println("Please press a valid step number")
				continue
			}

			currChapter = c.getChapter(currChapter.Options[i-1].Arc)
			break
		}
	}
}
