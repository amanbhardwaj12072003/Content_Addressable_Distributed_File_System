package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "dogsbestpicture"
	pathKey := CASPathTransformFunc(key)
	expectedFilename := "b519d9dda53fe3a9c3f1aca305b98a6f7d810a95"
	expectedPathName := "b519d/9dda5/3fe3a/9c3f1/aca30/5b98a/6f7d8/10a95"
	if pathKey.PathName != expectedPathName {
		t.Errorf("have: %s but want: %s", pathKey.PathName, expectedPathName)
	}
	if pathKey.FileName != expectedFilename {
		t.Errorf("have: %s but want: %s", pathKey.FileName, expectedFilename)
	}
}

func TestStore(t *testing.T) {
	s := newStore()
	id := generateID()
	defer teardown(t, s)
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("testcase_%d", i)
		data := []byte("some jpg bytes")

		if _, err := s.writeStream(id, key, bytes.NewReader(data)); err != nil {
			t.Error(err)
		}

		if ok := s.Has(id, key); !ok {
			t.Errorf("expected to have key %s", key)
		}

		_, r, err := s.Read(id, key)
		if err != nil {
			t.Error(err)
		}

		b, _ := ioutil.ReadAll(r)
		if string(b) != string(data) {
			t.Errorf("want: %s but have: %s", data, b)
		}

		if err := s.Delete(id, key); err != nil {
			t.Error(err)
		}

		if ok := s.Has(id, key); ok {
			t.Errorf("expected to NOT have key %s", key)
		}
	}
}

func newStore() *Store {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	return NewStore(opts)
}

func teardown(t *testing.T, s *Store) {
	if err := s.Clear(); err != nil {
		t.Error(err)
	}
}
