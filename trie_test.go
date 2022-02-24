package susliks

import (
    "testing"
)

func TestTrie(t *testing.T) {
    m := map[string]string{
        "abcd": "abcd",
        "a": "ab",
        "ax": "sss",
        "bcd": "fg",
    }

    trie := NewTrieFromMap(m)

    for k, v := range m {
        av, ok := trie.Lookup(k)
        if !ok {
            t.Fatalf("expected that key '%s' exist", k)
        }
        if av != v {
            t.Fatalf("expected value of key '%s' is '%s', actual is '%s'", k, v, av)
        }
    }

    trie.Delete("ax")
    trie.Delete("axx")

    value, ok := trie.Lookup("ax")
    if ok || value != "" {
        t.Fatalf("key 'ax' was deleted, but got %t, %s", ok, value)
    }
}
