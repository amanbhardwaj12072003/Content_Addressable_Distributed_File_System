package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

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

func TestPathTransformFunc(t *testing.T) {
	key := "dogsbestpicture"
	pathKey := CASPathTransformFunc(key)
	expectedOriginalKey := "b519d9dda53fe3a9c3f1aca305b98a6f7d810a95"
	expectedPathName := "b519d/9dda5/3fe3a/9c3f1/aca30/5b98a/6f7d8/10a95"
	if pathKey.PathName != expectedPathName {
		t.Errorf("have: %s but want: %s", pathKey.PathName, expectedPathName)
	}
	if pathKey.FileName != expectedOriginalKey {
		t.Errorf("have: %s but want: %s", pathKey.FileName, expectedOriginalKey)
	}
}

func TestStoreDelete(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	key := "myspecialpicture"
	data := []byte("some jpg bytes")

	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	if err := s.Delete(key); err != nil {
		t.Error(err)
	}
}

func TestStore(t *testing.T) {
	s := newStore()
	defer teardown(t, s)
	key := "myspecialpicture"
	data := []byte("some jpg bytes")

	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	if ok := s.Has(key); !ok {
		t.Errorf("expevted to have key %s", key)
	}

	r, err := s.Read(key)
	if err != nil {
		t.Error(err)
	}

	b, _ := ioutil.ReadAll(r)
	if string(b) != string(data) {
		t.Errorf("want: %s but have: %s", data, b)
	}

	s.Delete(key)
}
