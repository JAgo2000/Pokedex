package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("cache is nil")
	}

}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "",
			inputVal: []byte("val3"),
		},
	}

	for _, cas := range cases {
		cache.Add(cas.inputKey, []byte(cas.inputVal))
		actual, ok := cache.Get(cas.inputKey)
		if ok != true {
			t.Errorf("%v not found\n", cas.inputKey)
			continue
		}
		if string(actual) != string(cas.inputVal) {
			t.Errorf("%v doesent match %v\n", actual, cas.inputVal)
			continue
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
	}
	for _, cas := range cases {
		cache.Add(cas.inputKey, []byte(cas.inputVal))
	}
	_, ok := cache.Get(cases[0].inputKey)
	if !ok {
		t.Errorf("cache Value should exist <%v> but doesnt\n", string(cases[0].inputVal))
	}

	time.Sleep(interval + 1*time.Millisecond)
	_, ok = cache.Get(cases[0].inputKey)
	if ok {
		t.Errorf("doesnt reap old cache <%v>\n", string(cases[0].inputVal))
	}

}
