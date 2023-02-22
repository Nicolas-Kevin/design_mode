package design_mode

import "sync"

var (
	instance *Instance
	lock     sync.Once
)

type Instance struct {
	Name string
}

func GetInstance(name string) *Instance {
	if instance == nil {
		lock.Do(func() {
			instance = &Instance{
				Name: name,
			}
		})
	}
	return instance
}