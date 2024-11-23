package dBInMemory

type DB struct {
	data          map[int]string
	inTransaction bool
	transaction   map[int]string
}
