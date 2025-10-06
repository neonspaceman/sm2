package di

import (
	"card/internal/config"
	"slices"
	"strconv"
	"sync"
)

type Container int

// Deprecated: attempt to do DI
type DI struct {
	cfg        *config.Config
	mu         sync.Mutex
	cleanups   []func()
	containers map[Container]any
}

func NewDI(cfg *config.Config) *DI {
	return &DI{
		cfg:        cfg,
		containers: make(map[Container]any),
	}
}

func (di *DI) Close() {
	di.mu.Lock()
	defer di.mu.Unlock()

	slices.Reverse(di.cleanups)

	for _, cleanup := range di.cleanups {
		cleanup()
	}

	di.cleanups = nil
	clear(di.containers)
}

func (di *DI) addContainer(id Container, value any) {
	di.mu.Lock()
	defer di.mu.Unlock()

	if _, ok := di.containers[id]; ok {
		panic("duplicate container identification: " + strconv.Itoa(int(id)))
	}

	di.containers[id] = value
}

func (di *DI) getContainer(id Container) (any, bool) {
	di.mu.Lock()
	defer di.mu.Unlock()

	value, ok := di.containers[id]

	return value, ok
}

func (di *DI) addCleanup(cleanup func()) {
	di.mu.Lock()
	defer di.mu.Unlock()

	di.cleanups = append(di.cleanups, cleanup)
}

func createIfNil[T any](di *DI, id Container, creator func() (T, func(), error)) T {
	value, ok := di.getContainer(id)

	if ok {
		return value.(T)
	}

	value, cleanup, err := creator()

	if err != nil {
		panic(err)
	}

	if cleanup != nil {
		di.addCleanup(cleanup)
	}

	di.addContainer(id, value)

	return value.(T)
}
