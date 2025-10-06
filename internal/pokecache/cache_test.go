package pokecache

import (
    "testing"
    "time"
)

func TestCacheAddAndGet(t *testing.T) {
    cache := NewCache(2 * time.Second)
    key := "test-key"
    val := []byte("test-value")

    cache.Add(key, val)
    got, ok := cache.Get(key)
    if !ok {
        t.Fatal("expected key to be found")
    }
    if string(got) != string(val) {
        t.Fatalf("expected %s, got %s", val, got)
    }
}
