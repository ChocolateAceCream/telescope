package utils

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const (
	DEFAULT_TOTAL_SLOTS            = 3600
	DEFAULT_TIMEWHEEL_STEPDURATION = time.Microsecond
	DEFAULT_TIMEWHEEL_ERRORSIZE    = 1024 // 1.6kb
)

// timewheel running status
const (
	TIMEWHEEL_STATUS_IDEL = iota
	TIMEWHEEL_STATUS_RUNNING
	TIMEWHEEL_STATUS_END
)

type TimeWheel struct {
	name          string
	startTime     time.Time
	stepInterval  time.Duration
	totalSlot     int
	errChan       chan error
	status        int
	statusLock    sync.RWMutex
	quit          chan struct{}
	slots         []*slot
	taskMapper    map[uint64]*HandlerFunc
	currentID     *uint64
	cycleInterval time.Duration
	anchor        int // current slot index
}

type slot struct {
	mutex sync.RWMutex
	tasks *list.List
}
type HandlerFunc func() error

type task struct {
	id       uint64
	cycleNum int
}

// name string, totalSlots int, stepSize time.Duration, errChanBufferSize int

type Option func(*TimeWheel)

func WithTotalSlots(total int) Option {
	return func(tw *TimeWheel) {
		tw.totalSlot = total
	}
}

func WithInterval(size time.Duration) Option {
	return func(tw *TimeWheel) {
		tw.stepInterval = size
	}
}

func WithErrorChan(ch chan error) Option {
	return func(tw *TimeWheel) {
		tw.errChan = ch
	}
}

func WithName(name string) Option {
	return func(tw *TimeWheel) {
		tw.name = name
	}
}

func NewTimeWheel(opts ...Option) (tw *TimeWheel) {
	mapper := make(map[uint64]*HandlerFunc)
	currentID := uint64(0)
	errChan := make(chan error, DEFAULT_TIMEWHEEL_ERRORSIZE)
	tw = &TimeWheel{
		taskMapper: mapper,
		startTime:  time.Now(),
		currentID:  &currentID,
		totalSlot:  DEFAULT_TOTAL_SLOTS,
		quit:       make(chan struct{}),
		errChan:    errChan,
	}
	for _, opt := range opts {
		opt(tw)
	}
	slots := make([]*slot, tw.totalSlot)
	for i := range slots {
		slots[i] = &slot{
			tasks: list.New(),
		}
	}
	tw.slots = slots

	tw.cycleInterval = time.Duration(tw.totalSlot) * tw.stepInterval

	return
}

func (tw *TimeWheel) AddTask(delay time.Duration, handler HandlerFunc) (id uint64, err error) {
	if delay <= 0 {
		err = errors.New("delay must be greater than 0")
		return
	}
	slotIndex := int((delay % tw.cycleInterval) / tw.stepInterval)
	cycleNum := delay / tw.cycleInterval
	id = atomic.AddUint64(tw.currentID, 1)
	tw.slots[slotIndex].mutex.Lock()
	defer tw.slots[slotIndex].mutex.Unlock()
	tw.slots[slotIndex].tasks.PushBack(&task{
		id:       id,
		cycleNum: int(cycleNum),
	})
	tw.taskMapper[id] = &handler
	return
}

func (tw *TimeWheel) Run() (err error) {
	tw.statusLock.RLock()
	defer tw.statusLock.RUnlock()
	if tw.status == TIMEWHEEL_STATUS_RUNNING {
		err = errors.New("timewheel is already running")
		return
	}
	go func() {
		tw.statusLock.Lock()
		tw.status = TIMEWHEEL_STATUS_RUNNING
		tw.statusLock.Unlock()
		ticker := time.NewTicker(tw.stepInterval)
		for {
			select {
			case <-ticker.C:
				if tw.anchor >= tw.totalSlot {
					// new cycle
					tw.anchor = 0
				}
				tw.RunTask(tw.slots[tw.anchor])
				tw.anchor++
			case <-tw.quit:
				tw.status = TIMEWHEEL_STATUS_END
				return
			}

		}
	}()
	return
}

func (tw *TimeWheel) RunTask(s *slot) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for v := s.tasks.Front(); v != nil; v = v.Next() {
		n := v // copy the val, so it can used in go func
		if t := n.Value.(*task); t.cycleNum == 0 {
			go func() {
				defer func() {
					// clean up function
					s.mutex.Lock()
					s.tasks.Remove(n)
					s.mutex.Unlock()
				}()
				handler, ok := tw.taskMapper[t.id]
				if !ok {
					tw.errChan <- fmt.Errorf("task id %v not found", t.id)
					return
				}
				err := (*handler)()
				if err != nil {
					tw.errChan <- fmt.Errorf("task id %v run failed: %v", t.id, err)
					return
				}
			}()
		} else {
			t.cycleNum--
		}

	}
}

func (tw *TimeWheel) Stop() (err error) {
	tw.statusLock.RLock()
	defer tw.statusLock.RUnlock()
	if tw.status != TIMEWHEEL_STATUS_RUNNING {
		err = errors.New("timewheel is not running")
		return
	}
	close(tw.quit)
	return
}
