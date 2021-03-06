package databases

import (
	"errors"

	"github.com/go-in-memory-db/queue"
)

// CreateQueue Function
func (db *Database) CreateQueue(k string) (*queue.Queue, error) {
	if queue, ok := db.queue[k]; ok {
		return queue, errors.New("Hash Table Exists")
	}
	db.queue[k] = queue.NewQueue()
	return db.queue[k], nil
}

// GetQueue Function
func (db *Database) GetQueue(k string) (*queue.Queue, error) {
	if _, ok := db.queue[k]; ok {
		return db.queue[k], nil
	}
	return nil, errors.New("not found")
}

// DelQueue Function
func (db *Database) DelQueue(k string) {
	delete(db.queue, k)
}
