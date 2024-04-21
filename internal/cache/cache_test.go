package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	interval := 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
        {
            key: "https://example.com",
            val: []byte("data"),
        },
        {
            key: "https://example.com/foo",
            val: []byte("moredata"),
        },
        {
            key: "https://example.com/bar",
            val: []byte("somemoredata"),
        },
    }

    for i, c := range cases {
        t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
            cache := NewCache(interval)
            cache.Add(c.key, c.val)
            val, ok := cache.Get(c.key)
            if !ok {
                t.Errorf("expected to find key")
                return
            }
            if string(val) != string(c.val) {
                t.Errorf("expected to find value")
            }
        })
    }
}

func TestReapLoop(t *testing.T) {
    baseTime := 5 * time.Millisecond
    waitTime := baseTime + 5 * time.Millisecond
    cache := NewCache(baseTime)
    cache.Add("https://example.com", []byte("data"))

    _, ok := cache.Get("https://example.com")
    if !ok {
        t.Errorf("expected to find key")
        return
    }

    time.Sleep(waitTime)

    _, ok = cache.Get("https://example.com")
    if ok {
        t.Errorf("expected to not find key")
        return
    }
}
