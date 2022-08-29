package story

import (
	"encoding/json"
	"io/ioutil"
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
	f, err := ioutil.ReadFile(filePath)

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
