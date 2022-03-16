package event

import (
	"context"
	"fmt"

	"git.internal.yunify.com/manage/common/log"
)

const asyncName = "asyncer"
const CountWorker = 20

var manager *EventManager

func init() {
	ctx := context.Background()
	manager = &EventManager{
		ctx:           ctx,
		countWorker:   CountWorker,
		AsyncWorkers:  make([]*AsyncWorker, 0), //NewAsyncWorker(ctx, asyncName),
		Subscriptions: NewSubscription(),
	}
	for index := 0; index < manager.countWorker; index++ {
		manager.AsyncWorkers = append(manager.AsyncWorkers, NewAsyncWorker(ctx, fmt.Sprintf("%s-%d", asyncName, index)))
	}
}

type EventManager struct {
	ctx           context.Context
	countWorker   int
	AsyncWorkers  []*AsyncWorker
	Subscriptions *Subscription
}

func GetEventManager() *EventManager {
	return manager
}

func (this *EventManager) Subscrib(s Subscriber) {
	this.Subscriptions.Subscrib(s)
}

func (this *EventManager) Unsubscrib(s Subscriber) {
	this.Subscriptions.Unsubscrib(s)
}

func (this *EventManager) Run() {

	log.InfoWithFields(eventSubscriptionLogTitle, log.Fields{
		"desc":    "start eventer....",
		"Workers": this.countWorker,
	})
	for _, worker := range this.AsyncWorkers {
		go worker.Start()
	}
}

func (this *EventManager) Stop() {
	log.InfoWithFields(eventSubscriptionLogTitle, log.Fields{
		"desc":    "stop eventer....",
		"Workers": this.countWorker,
	})
	for _, worker := range this.AsyncWorkers {
		go worker.Stop()
	}
}

func (this *EventManager) Send(event IEvent) {
	worker := this.loadWorker()
	this.Subscriptions.Send(event, worker)
}

func (this *EventManager) loadWorker() *AsyncWorker {
	var index, weight int = 0, this.AsyncWorkers[0].Len()
	for indx, worker := range this.AsyncWorkers {
		if wt := worker.Len(); weight < wt {
			weight = wt
			index = indx
		}
	}
	return this.AsyncWorkers[index]
}
