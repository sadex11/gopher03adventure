package story

import (
	"encoding/json"
	"os"
)

type StoryConfig map[string]Chapter

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func readStoryMap(filePath string) []byte {
	f, err := os.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	return f
}

func CreateStories(filePath string) StoryConfig {
	storyContent := readStoryMap(filePath)

	storyMap := make(map[string]Chapter)

	err := json.Unmarshal(storyContent, &storyMap)

	if err != nil {
		panic(err)
	}

	return storyMap
}
