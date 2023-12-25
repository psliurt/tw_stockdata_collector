package env

import (
	"sync"
)

var envInstance *Environment
var createEnvironmentOnce sync.Once

type Environment struct {
}

func Initialize() {
	createEnvironmentOnce.Do(func() {
		envInstance = createEnvironment()
	})
}

func Instance() *Environment {
	Initialize()
	return envInstance
}

func createEnvironment() *Environment {
	setUpZap()
	return &Environment{}
}
