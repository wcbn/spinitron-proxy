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
		"/images/Persona/16/65/166599-img_profile.225x225.jpg?v=123",
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

func TestIsNotCollectionPath(t *testing.T) {

	known := []string{
		"/api/personas/1",
		"/images/Persona/16/65/166599-img_profile.225x225.jpg?v=123",
	}

	for _, name := range known {
		result := !IsCollectionPath(name)
		if !result {
			t.Errorf("IsCollectionPath(%s) = %t; want true", name, result)
		}
	}
}

func TestGetCollectionName(t *testing.T) {
	s := []string{
		"api/foo",
		"foo",
		"/api/foo",
		"/foo",
		"/foo/",
		"/api/foo/",
		"/foo/",
		"/api/foo/?bar=baz",
		"/api/foo?bar=baz",
	}

	for _, name := range s {
		result := GetCollectionName(name)
		if result != "foo" {
			t.Errorf("GetCollectionName(%s) = %s; want foo", name, result)
		}
	}
}
