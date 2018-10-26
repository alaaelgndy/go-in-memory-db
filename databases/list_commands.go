package databases

import (
	"errors"

	"github.com/go-in-memory-db/linkedlist"
)

func (db *Database) GetList(k string) (*linkedlist.LinkedList, error) {
	if _, ok := db.dataList[k]; ok {
		return db.dataList[k], nil
	}

	return nil, errors.New("not found")

}

func (db *Database) CreateList(k string) (*linkedlist.LinkedList, error) {
	if v, ok := db.dataList[k]; ok {
		return v, errors.New("List Exists")
	}
	db.dataList[k] = linkedlist.NewList()
	return db.dataList[k], nil
}

func (db *Database) DelList(k string) {
	delete(db.dataList, k)
}
