package lru

import "container/list"

type LRUCache struct {
	capacity   int
	cache      *list.List
	keyToValue map[int]*list.Element
}

type entry struct {
	key   int
	value int
}

func Contructor(capacity int) LRUCache {
	return LRUCache{
		capacity:   capacity,
		cache:      list.New(),
		keyToValue: make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {

	if el, ok := this.keyToValue[key]; ok {
		this.cache.MoveToFront(el)
		return el.Value.(*entry).value
	}

	return -1
}

func (this *LRUCache) Put(key, value int) {

	if el, ok := this.keyToValue[key]; ok {
		this.cache.MoveToFront(el)
		el.Value.(*entry).value = value
	} else {
		if this.cache.Len() == this.capacity {
			oldest := this.cache.Back()
			if oldest != nil {

				this.cache.Remove(oldest)
				delete(this.keyToValue, oldest.Value.(*entry).key)
			}
		}

		el := this.cache.PushFront(&entry{key: key, value: value})
		this.keyToValue[key] = el
	}
}
