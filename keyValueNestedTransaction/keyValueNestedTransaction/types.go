package keyValueNestedTransaction

import "sync"

type KVStore struct {
	hm               map[string]string
	transactionStack []map[string]string
	delMark          string
	mutex            sync.Mutex
}

func NewKVStore() *KVStore {
	return &KVStore{
		hm:               make(map[string]string),
		transactionStack: []map[string]string{},
		delMark:          "__DEL__",
	}
}
