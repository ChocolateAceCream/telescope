package dataInitializer

import (
	"context"
	"sort"
	"sync"

	"github.com/ChocolateAceCream/telescope/backend/singleton"
)

const (
	internalOrder = 10
	externalOrder = 100
)

type Initializer interface {
	Name() string
	Init(ctx context.Context) (next context.Context, err error)
	VerifyData(ctx context.Context) bool
}

type orderedInitializer struct {
	order int
	Initializer
}

func init() {
	once.Do(func() {
		cache = make(map[string]*orderedInitializer)
	})
}

type initializerGroup []*orderedInitializer

var (
	once             sync.Once
	dataInitializers initializerGroup
	cache            map[string]*orderedInitializer
)

func Register(order int, initializer Initializer) {
	name := initializer.Name()
	oi := &orderedInitializer{
		order:       order,
		Initializer: initializer,
	}
	dataInitializers = append(dataInitializers, oi)
	cache[name] = oi
}

func InitData() (err error) {
	ctx := context.TODO()
	if len(dataInitializers) == 0 {
		return
	}
	sort.Sort(&dataInitializers)
	next, cancel := context.WithCancel(ctx)
	defer cancel()
	for _, initializer := range dataInitializers {
		if initializer.VerifyData(next) {
			// already initialized, continue
			continue
		}
		next, err = initializer.Init(next)
		if err != nil {
			return
		}
	}
	singleton.Logger.Info("init data success")
	return
}

func (ig initializerGroup) Len() int {
	return len(ig)
}

func (ig initializerGroup) Less(i, j int) bool {
	return ig[i].order < ig[j].order
}

func (ig initializerGroup) Swap(i, j int) {
	ig[i], ig[j] = ig[j], ig[i]
}
