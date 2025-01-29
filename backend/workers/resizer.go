/*
* @fileName resizer.go
* @author Di Sheng
* @date 2024/09/18 21:58:42
* @description: used to resize image
 */

package workers

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/ChocolateAceCream/telescope/backend/lib"
	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

type resizer struct {
	Started      bool
	Consumer     sarama.ConsumerGroup
	quit         chan struct{}
	WorkersCount int
	wg           sync.WaitGroup
}

type Message struct {
	Key int32
	Val string
}

func (r *resizer) Name() string {
	return "resizer"
}

func (r *resizer) Count() int {
	return r.WorkersCount
}

func (r *resizer) FeedJobs(jobs chan Message) {
	brokers := []string{"localhost:29092", "localhost:39092", "localhost:49092"}
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumerGroup, err := sarama.NewConsumerGroup(brokers, "resizer", config)
	if err != nil {
		singleton.Logger.Error("Failed to create kafka consumer group", zap.String("error", err.Error()))
		return
	}
	r.Consumer = consumerGroup
	r.quit = make(chan struct{})
	fmt.Println("----r.quit:", r.quit)

	ctx := context.Background()
	for {
		select {
		case <-r.quit:
			singleton.Logger.Info("Stopping FeedJobs...")
			r.wg.Done()
			return
		default:
			handler := &consumerGroupHandler{jobs: jobs}
			err := r.Consumer.Consume(ctx, []string{"admin1"}, handler)
			if err != nil {
				singleton.Logger.Error("Error from consumer", zap.String("error", err.Error()))
			}
		}
	}
}

func (r *resizer) Start(ctx context.Context) (next context.Context, err error) {
	next = ctx
	r.WorkersCount = singleton.Config.Workers.Resizer.Count
	fmt.Println("Starting resizer with", r.WorkersCount, "workers")
	jobs := make(chan Message, singleton.Config.Workers.Resizer.QueueSize)
	for w := 1; w <= r.WorkersCount; w++ {
		go worker(w, jobs)
	}
	r.Started = true
	go r.FeedJobs(jobs)
	return
}

func (r *resizer) Stop() {
	if r.quit != nil && r.Started {
		r.wg.Add(1)
		go close(r.quit)
		r.wg.Wait()
		fmt.Println("Closing consumer")
		err := r.Consumer.Close()
		fmt.Println("Consumer closed, err:", err)

		r.Started = false
		singleton.Logger.Info("resizer stopped")
	}
}

func (r *resizer) Restart(ctx context.Context) (err error) {
	r.Stop()
	_, err = r.Start(ctx) // Start it again
	if err != nil {
		singleton.Logger.Error("Failed to restart resizer", zap.String("error", err.Error()))
		return
	}
	singleton.Logger.Info("Resizer restarted.")
	return
}

func (r *resizer) Verified(ctx context.Context) bool {
	return r.Started
}

func init() {
	Register(&resizer{})
}

func worker(id int, jobs <-chan Message) {
	for j := range jobs {
		time.Sleep(5 * time.Second)
		singleton.Logger.Info("---------working on jobs------------", zap.Int("id", id), zap.Any("job", j))
		ch := lib.GetActiveSSE(j.Key)
		ch <- string(j.Val)
	}
	singleton.Logger.Info("Worker exiting", zap.Int("id", id))
}

type consumerGroupHandler struct {
	jobs chan Message
}

func (h *consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (h *consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

// ConsumeClaim processes messages from the assigned partition
func (h *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// consume message from the topic
	for msg := range claim.Messages() {
		key, _ := strconv.Atoi(string(msg.Key))
		m := Message{Key: int32(key), Val: string(msg.Value)}
		h.jobs <- m

		// Mark message as processed
		session.MarkMessage(msg, "")
	}
	return nil
}
