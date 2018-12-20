package hash_table

import (
	"log"
	"testing"
)

func TestSimple(t *testing.T) {
	table := NewHashTable(3)

	log.Println(table.GetKeys())
	table.Set("test", 0)
	table.Set("test2", 1)
	log.Println(table.GetKeys())
	table.Set("test3", "hopa")
	table.ForEach(func(key, value interface{}) {
		log.Println("key:", key, ", value:", value)
	})

	table.Delete("test2")
	log.Println(table.GetKeys())
}
