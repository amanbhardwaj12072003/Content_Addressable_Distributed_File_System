package main

import (
	"bytes"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "dogsbestpicture"
	pathKey := CASPathTransformFunc(key)
	expectedOriginalKey := "b519d9dda53fe3a9c3f1aca305b98a6f7d8/10a95"
	expectedPathName := "b519d/9dda5/3fe3a/9c3f1/aca30/5b98a/6f7d8/10a95"
	if pathKey.PathName != expectedPathName {
		t.Errorf("have: %s but want: %s", pathKey.PathName, expectedPathName)
	}
	if pathKey.FileName != expectedOriginalKey {
		t.Errorf("have: %s but want: %s", pathKey.FileName, expectedOriginalKey)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	key := "myspecialpicture"

	data := bytes.NewReader([]byte("some jpg bytes"))
	if err := s.writeStream(key, data); err != nil {
		t.Error(err)
	}
}
