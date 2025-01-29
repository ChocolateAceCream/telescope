package workers

import (
	"context"
	"errors"
	"sync"

	"github.com/ChocolateAceCream/telescope/backend/singleton"
)

var (
	once                sync.Once
	workersInitializers initializerGroup
	cache               map[string]*workerPool
)

type workerPool struct {
	Workers
}

type initializerGroup []*workerPool

type Workers interface {
	Name() string
	Count() int
	Start(ctx context.Context) (next context.Context, err error)
	Verified(ctx context.Context) bool
	Restart(ctx context.Context) (err error)
}

func init() {
	once.Do(func() {
		cache = make(map[string]*workerPool)
	})
}

func Register(workers Workers) {
	name := workers.Name()
	oi := &workerPool{
		Workers: workers,
	}
	workersInitializers = append(workersInitializers, oi)
	cache[name] = oi
}

func StartWorkerPool() (err error) {
	ctx := context.TODO()
	if len(workersInitializers) == 0 {
		return
	}
	// sort.Sort(&workersInitializers)
	next, cancel := context.WithCancel(ctx)
	defer cancel()
	for _, initializer := range workersInitializers {
		if initializer.Verified(next) {
			// already initialized, continue
			continue
		}
		next, err = initializer.Start(next)
		if err != nil {
			return
		}
	}
	singleton.Logger.Info("all worker pools start success")
	return
}

func GetWorkerCountByWorkerPoolType(name string) (count int, ok bool) {
	wp, ok := cache[name]
	if !ok {
		return
	}
	return wp.Count(), true
}

func RestartWorkerPoolByPoolName(name string) (err error) {
	wp, ok := cache[name]
	if !ok {
		err = errors.New("worker pool not found")
		return
	}
	err = wp.Workers.Restart(context.TODO())
	return
}
