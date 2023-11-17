package pokecache

import "testing"

func TestCreateCache(t *testing.T) {
	cache := NewCache()
	if cache.cache == nil {
		t.Error("cache is nil")
	}

}

func TestAddGetCache(t *testing.T) {
	cache := NewCache()

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
			t.Error("key1 not found")
			continue
		}
		if string(actual) != string(cas.inputVal) {
			t.Errorf("%v doesent match %v\n", actual, cas.inputVal)
			continue
		}
	}
}
