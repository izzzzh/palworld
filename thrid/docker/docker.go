package docker

import "sync"

var ClientPool Client

func InitClient() {
	once := &sync.Once{}
	once.Do(func() {
		client, err := NewClient()
		if err != nil {
			return
		}
		ClientPool = client
	})
	return
}
