package event

import (
	"context"

	"git.internal.yunify.com/manage/common/log"
)

const eventAsyncerLogTitle = "[EventAsyncer]"

type WorkerItem struct {
	Ev         IEvent
	Subscriber Subscriber
}

// TODO
// just a simple go routine pool?
type AsyncWorker struct {
	Ctx    context.Context
	Cancel context.CancelFunc
	Name   string
	// better to use queue
	ItemChan chan WorkerItem
	Running  bool
}

func NewAsyncWorker(ctx context.Context, name string) *AsyncWorker {
	as := new(AsyncWorker)
	ctx, cancel := context.WithCancel(ctx)
	as.Ctx = ctx
	as.Cancel = cancel
	as.ItemChan = make(chan WorkerItem, 1024)
	as.Name = name

	//go as.Start()
	return as
}

func (this *AsyncWorker) AsyncHandle(event IEvent, subscriber Subscriber) {
	item := WorkerItem{
		Ev:         event,
		Subscriber: subscriber,
	}
	this.ItemChan <- item
}

func (this *AsyncWorker) Len() int {
	return len(this.ItemChan)
}

func (this *AsyncWorker) Start() {
	defer func() {
		err := recover()
		if err != nil {
			log.DebugWithFields(eventAsyncerLogTitle, log.Fields{
				"desc": "event async worker error",
			})
		}
	}()
	this.Running = true
	for {
		select {
		case <-this.Ctx.Done():
			return
		case item := <-this.ItemChan:
			this.handleEvent(item)
		}
	}
}

func (this *AsyncWorker) handleEvent(item WorkerItem) {
	item.Subscriber.OnEvent(item.Ev)
}

func (this *AsyncWorker) Stop() {
	if this.Running {
		if this.Len() > 0 {
			log.WarnWithFields(eventSubscriptionLogTitle, log.Fields{
				"desc":  "stopping eventer....",
				"tasks": this.Len(),
			})
		}
		this.Cancel()
		close(this.ItemChan)
	}
}
