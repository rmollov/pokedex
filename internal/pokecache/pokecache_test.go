package pokecache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {

	cache := NewCache(5 * time.Second)

	cache.Add("hello", []byte("some test data"))
	actual, ok := cache.Get("hello")

	if !ok {
		t.Errorf("missing key")
	}

	if string(actual) != "some test data" {
		t.Errorf("data differs")
	}

}

func TestExpire(t *testing.T) {

	cache := NewCache(50 * time.Millisecond)

	cache.Add("hello", []byte("some test data"))

	time.Sleep(1 * time.Second)
	_, ok := cache.Get("hello")

	if ok {
		t.Errorf("key still present")
	}

}
