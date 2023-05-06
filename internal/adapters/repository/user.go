package repository

import "sync"

type UserRepository struct{
	mux  *sync.RWMutex
	
}