package susliks

import (
    "unicode/utf8"
)

type Trie[T any] struct {
    childs   map[rune]*Trie[T]
    value    Maybe[T]
}

func NewTrie[T any]() *Trie[T] {
    return &Trie[T]{}
}

func NewTrieFromMap[T any](m map[string]T) *Trie[T] {
    trie := NewTrie[T]()

    for k, v := range m {
        trie.Insert(k, v)
    }

    return trie
}

func (trie *Trie[T]) Insert(key string, value T) {
    for _, k := range key {
        if len(trie.childs) == 0 {
            trie.childs = make(map[rune]*Trie[T])
        }
        if _, ok := trie.childs[k]; !ok {
            trie.childs[k] = NewTrie[T]()
        }
        trie = trie.childs[k]
    }

    trie.value = Just(value)
}

func (trie *Trie[T]) Delete(key string) {
    trie.del(key)
}

func (trie *Trie[T]) del(key string) bool {
    if key == "" {
        trie.value = Nothing[T]()
    } else {
        k, size := utf8.DecodeRuneInString(key)
        if _, ok := trie.childs[k]; !ok {
            return false
        }

        if trie.childs[k].del(key[size:]) {
            delete(trie.childs, k)
        }
    }

    return len(trie.childs) == 0
}

func (trie *Trie[T]) Lookup(key string) (T, bool) {
    var zero T

    for _, k := range key {
        if _, ok := trie.childs[k]; !ok {
            return zero, false
        }
        trie = trie.childs[k]
    }

    return trie.value.Unwrap()
}

/*
func (trie *Trie[T]) FindAllByPrefix(key string) []T {
    for _, k := range key {
        if _, ok := trie.childs[k]; !ok {
            return []T
        }
        trie = trie.childs[k]
    }

    var res []T

    convertTrieToSlice(trie, &res)
}

func convertTrieToSlice[T any](trie *Trie[T], res *[]T) {
    for _, k := range trie.
}

*/
