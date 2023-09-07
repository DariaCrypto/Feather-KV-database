package inmemory

import (
    "errors"
)

type HashMapCollection struct {
    name    string
    hashMap map[string]*HashMap
}

type SortedSetCollection struct {
    name      string
    sortedSet map[string]*SortedSet
}

func newHashMapCollection(name string) *HashMapCollection {
    hms := make(map[string]*HashMap)
    return &HashMapCollection{name, hms}
}

func newSortedSetCollection(name string) *SortedSetCollection {
    ss := make(map[string]*SortedSet)
    return &SortedSetCollection{name, ss}
}

func (hm *HashMapCollection) GetHashMap(name string) (*HashMap, error) {
    v, ok := hm.hashMap[name]
    if !ok {
        return nil, errors.New("collections: storage not found")
    }
    return v, nil
}

func (ss *SortedSetCollection) GetSortedSet(name string) (*SortedSet, error) {
    v, ok := ss.sortedSet[name]
    if !ok {
        return nil, errors.New("collections: storage not found")
    }
    return v, nil
}

func (hm *HashMapCollection) DeleteHashMap(name string) error {
    _, ok := hm.hashMap[name]
    if !ok {
        return errors.New("collections: storage not found")
    }
    delete(hm.hashMap, name)
    return nil
}

func (ss *SortedSetCollection) DeleteSortedSet(name string) error {
    _, ok := ss.sortedSet[name]
    if !ok {
        return errors.New("collections: storage not found")
    }
    delete(ss.sortedSet, name)
    return nil
}

func (hm *HashMapCollection) CreateHashMap(name string) (*HashMap, error) {
    if _, ok := hm.hashMap[name]; ok {
        return nil, errors.New("collections: storage already exists")
    }
    hm.hashMap[name] = NewHashMap(name)
    return hm.hashMap[name], nil
}

func (ss *SortedSetCollection) CreateSortedSet(name string) (*SortedSet, error) {
    if _, ok := ss.sortedSet[name]; ok {
        return nil, errors.New("collections: storage already exists")
    }
    ss.sortedSet[name] = NewSortedSet(name)
    return ss.sortedSet[name], nil
}
