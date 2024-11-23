package dBInMemory

import "errors"

func (db *DB) Get(id int) (string, bool) {
	if db.inTransaction {
		if value, exists := db.transaction[id]; exists {
			return value, exists
		}
	}
	value, exists := db.data[id]
	return value, exists
}

func (db *DB) Set(id int, val string) {
	if db.inTransaction {
		db.transaction[id] = val
	} else {
		db.data[id] = val
	}
}

func (db *DB) DeleteByID(id int) {

	if db.inTransaction {
		delete(db.transaction, id)
	} else {
		delete(db.data, id)
	}
}

func (db *DB) DeleteByValue(val string) {
	if db.inTransaction {
		for id, v := range db.transaction {
			if v == val {
				delete(db.transaction, id)
			}
		}
	} else {
		for id, v := range db.data {
			if v == val {
				delete(db.data, id)
			}
		}
	}
}

func (db *DB) Begin() error {
	if db.inTransaction {
		return errors.New("already in transaction")
	}
	db.inTransaction = true
	db.transaction = make(map[int]string)
	return nil
}

func (db *DB) Commit() error {
	if !db.inTransaction {
		return errors.New("not in transaction to commit")
	}
	for k, val := range db.transaction {
		db.data[k] = val
	}
	db.transaction = nil
	db.inTransaction = false
	return nil
}

func (db *DB) Rollback() error {
	if !db.inTransaction {
		return errors.New("not in transaction to Rollback")
	}
	db.transaction = nil
	db.inTransaction = false
	return nil
}

func NewDB() *DB {
	data := make(map[int]string)
	transaction := make(map[int]string)

	return &DB{
		data:        data,
		transaction: transaction,
	}
}
