package handler

import "testing"

func TestGetStoryPathEmpty(t *testing.T) {
	res := getStoryPath("")

	if res != "intro" {
		t.Error("Empty string return invalid path '", res, "', expected 'intro'")
	}
}

func TestGetStoryPathBase(t *testing.T) {
	res := getStoryPath("/")

	if res != "intro" {
		t.Error("Base string path return invalid path '", res, "', expected 'intro'")
	}
}

func TestGetStoryPath(t *testing.T) {
	res := getStoryPath("/test")

	if res != "test" {
		t.Error("String path return invalid path'", res, "', expected 'test'")
	}
}
