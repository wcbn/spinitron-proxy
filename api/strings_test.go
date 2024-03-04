package api

import (
	"testing"
)

func TestIsResourcePath(t *testing.T) {

	known := []string{
		"/api/personas/1",
		"/api/shows/2",
		"/api/playlists/3",
		"/api/spins/4",
	}

	for _, name := range known {
		result := IsResourcePath(name)
		if !result {
			t.Errorf("IsResourcePath(%s) = %t; want true", name, result)
		}
	}
}

func TestIsCollectionPath(t *testing.T) {

	known := []string{
		"/api/personas",
		"/api/shows",
		"/api/playlists",
		"/api/spins",
	}

	for _, name := range known {
		result := IsCollectionPath(name)
		if !result {
			t.Errorf("IsCollectionPath(%s) = %t; want true", name, result)
		}
	}
}
