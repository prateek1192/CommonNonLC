package keyValueNestedTransaction

import "fmt"

func (k *KVStore) Get(key string) (string, bool) {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	for i := len(k.transactionStack) - 1; i >= 0; i-- {
		if val, exists := k.transactionStack[i][key]; exists {
			if val == k.delMark {
				return "", false
			}
			return val, true
		}
	}
	val, exists := k.hm[key]
	return val, exists
}

func (k *KVStore) Set(key, value string) {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	if len(k.transactionStack) > 0 {
		k.transactionStack[len(k.transactionStack)-1][key] = value
	} else {
		k.hm[key] = value
	}
}

func (k *KVStore) Delete(key string) {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	if len(k.transactionStack) > 0 {
		top := len(k.transactionStack) - 1
		k.transactionStack[top][key] = k.delMark
	} else {
		delete(k.hm, key)
	}
}

func (k *KVStore) Begin() {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	k.transactionStack = append(k.transactionStack, make(map[string]string))
}

func (k *KVStore) Commit() {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	if len(k.transactionStack) == 0 {
		fmt.Printf("Nothing to commit")
		return
	}
	top := len(k.transactionStack) - 1
	for key, val := range k.transactionStack[top] {
		if val == k.delMark {
			delete(k.hm, key)
		} else {
			k.hm[key] = val
		}
	}
	k.transactionStack = k.transactionStack[:top]
}

func (k *KVStore) Rollback() {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	if len(k.transactionStack) == 0 {
		fmt.Printf("No transaction to rollback \n")
		return
	}
	top := len(k.transactionStack) - 1
	k.transactionStack = k.transactionStack[:top]
}
