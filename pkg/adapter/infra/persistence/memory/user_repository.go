package memory

import "sync"

type UserRepository struct {
	sync.RWMutex
}
