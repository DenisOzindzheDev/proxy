package cache

import (
	"container/list"
	"rproxy/pkg/logger"
)

type Item struct {
	Key   string
	Value interface{}
}

type LRU struct {
	capacity int
	items    map[string]*list.Element
	queue    *list.List
}

func NewLru(capacity int) *LRU {
	return &LRU{capacity: capacity, items: make(map[string]*list.Element), queue: list.New()}
}

func (c *LRU) SetKey(key string, value interface{}) error {
	if element, ok := c.items[key]; ok {
		c.queue.MoveToFront(element)
		element.Value.(*Item).Value = value
		return nil
	}
	if c.queue.Len() == c.capacity {
		logger.Warnf("[CACHE]capacity %d exceeded", c.capacity) //not sure about this
		c.purge()
	}

	item := &Item{
		Key:   key,
		Value: value,
	}
	element := c.queue.PushFront(item)
	c.items[item.Key] = element

	return nil
}

func (c *LRU) purge() {
	if element := c.queue.Back(); element != nil {
		item := c.queue.Remove(element).(*Item)
		delete(c.items, item.Key)
	}
}
func (c *LRU) GetKey(key string) interface{} {
	if _, ok := c.items[key]; !ok {
		return nil
	}
	c.queue.MoveToFront(c.items[key])
	return c.items[key].Value.(*Item).Value

}
