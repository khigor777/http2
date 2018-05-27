package users

import "sync"

type Users struct {
	mu sync.RWMutex
	Users []*Account
}

type Account struct {
	Name, Email string
}

func (u *Users) Add (account *Account)  {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.Users = append(u.Users, account)
}

func (u *Users) GetAll()[]*Account  {
	u.mu.RLock()
	defer u.mu.Unlock()
	return u.Users
}


