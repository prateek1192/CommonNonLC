A thread-safe in-memory cache implementation in Go with a simple eviction mechanism. This package supports concurrent access and ensures data integrity using read-write locks.

## Features

- **Thread-Safe**: Uses `sync.RWMutex` to ensure safe concurrent access.
- **Eviction Policy**: Implements a basic eviction mechanism where the least recently used (LRU) item is removed when the cache reaches its capacity.
- **Simple API**: Provides `Set`, `Get`, and automatic eviction functionality.