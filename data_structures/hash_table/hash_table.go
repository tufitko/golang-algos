package hash_table

import (
	"fmt"
	"github.com/tufitko/golang-algos/data_structures/linked_list"
	"log"
	"sync"
)

type HashTableEntry struct {
	key, value interface{}
}

type HashTable struct {
	array []*linked_list.LinkedList
}

func NewHashTable(size int) *HashTable {
	array := make([]*linked_list.LinkedList, size)
	for i := 0; i < size; i++ {
		array[i] = linked_list.NewLinkedList()
	}

	return &HashTable{array: array}
}

func (table *HashTable) hash(value interface{}) int {
	stringView := fmt.Sprintf("%+v\n", value)

	sumBytes := 0
	for _, b := range []byte(stringView) {
		sumBytes += int(b)
	}

	return sumBytes % len(table.array)
}

func (table *HashTable) Set(key, value interface{}) {
	table.Delete(key)
	hash := table.hash(key)
	table.array[hash].Append(HashTableEntry{key: key, value: value})
}

func (table *HashTable) Get(key interface{}) interface{} {
	hash := table.hash(key)

	for i := 0; i < table.array[hash].Len(); i++ {
		entry := table.array[hash].Get(i).(HashTableEntry)
		if entry.key == key {
			return entry.value
		}
	}

	return nil
}

func (table *HashTable) HashKey(key interface{}) bool {
	hash := table.hash(key)

	for i := 0; i < table.array[hash].Len(); i++ {
		entry := table.array[hash].Get(i).(HashTableEntry)
		if entry.key == key {
			return true
		}
	}

	return false
}

func (table *HashTable) Delete(key interface{}) {
	hash := table.hash(key)

	for i := 0; i < table.array[hash].Len(); i++ {
		entry := table.array[hash].Get(i).(HashTableEntry)
		if entry.key == key {
			table.array[hash].Delete(i)
		}
	}
}

//TODO GetKeys with this func dont work normally
func (table *HashTable) ForEachParallel(fn func(key, value interface{})) {
	wg := sync.WaitGroup{}

	for _, list := range table.array {
		wg.Add(list.Len())
		go list.ForEach(func(value interface{}) {
			log.Println(wg)
			entry := value.(HashTableEntry)
			fn(entry.key, entry.value)
			wg.Done()
		})
	}

	wg.Wait()
}

func (table *HashTable) ForEach(fn func(key, value interface{})) {
	for _, list := range table.array {
		list.ForEach(func(value interface{}) {
			entry := value.(HashTableEntry)
			fn(entry.key, entry.value)
		})
	}
}

func (table *HashTable) GetKeys() []interface{} {
	keys := make([]interface{}, 0)

	table.ForEach(func(key, _ interface{}) {
		keys = append(keys, key)
	})

	return keys
}
