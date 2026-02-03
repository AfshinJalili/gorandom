package commands

import "sync"

func resetSourcesFetchOnce() {
	sourcesFetchOnce = sync.Once{}
}
