package event

import (
	"sync"

	"git.internal.yunify.com/manage/common/log"
)

const eventSubscriptionLogTitle = "[EventSub]"

type Subscription struct {
	Subscriptions map[string]map[string][]Subscriber //map[name]map[rule_id][]subscriber
	Lock          sync.RWMutex
}

func NewSubscription() *Subscription {
	return &Subscription{
		Subscriptions: make(map[string]map[string][]Subscriber),
	}
}
func (s *Subscription) Subscrib(subscriber Subscriber) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	name := subscriber.Name()
	id := subscriber.Id()
	if _, has := s.Subscriptions[name]; !has {
		s.Subscriptions[name] = make(map[string][]Subscriber)
	}
	subs := s.Subscriptions[name]
	if _, has := subs[id]; !has {
		subs[id] = make([]Subscriber, 0)
	}
	log.InfoWithFields(eventSubscriptionLogTitle, log.Fields{
		"desc":  "subscrib",
		"name:": subscriber.Name(),
		"id:":   id,
	})
	subs[id] = append(subs[id], subscriber)
}

func (s *Subscription) Unsubscrib(subscriber Subscriber) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	name := subscriber.Name()
	id := subscriber.Id()
	if _, has := s.Subscriptions[name]; !has {
		return
	}
	if _, has := s.Subscriptions[name][id]; !has {
		return
	}
	ss := s.Subscriptions[name][id]
	for index, sb := range ss {
		if sb.Id() == subscriber.Id() {
			ss = append(ss[0:index], ss[index+1:]...)
		}
	}
	log.InfoWithFields(eventSubscriptionLogTitle, log.Fields{
		"desc":  "unsubscrib",
		"name:": subscriber.Name(),
		"id:":   id,
	})
}

func (s *Subscription) getSubscribers(name, id string) []Subscriber {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	subscribers := make([]Subscriber, 0)
	if _, has := s.Subscriptions[name]; has {
		if sbs, has := s.Subscriptions[name][id]; has {
			subscribers = append(subscribers, sbs...)
		}
		//default.
		if sbs, has := s.Subscriptions[name]["default"]; has {
			subscribers = append(subscribers, sbs...)
		}
	}
	return subscribers
}

func (s *Subscription) Send(e IEvent, asyncer *AsyncWorker) {
	// log.InfoWithFields(eventSubscriptionLogTitle, log.Fields{
	// 	"desc":  "send message",
	// 	"name:": e.Name(),
	// 	"id:":   e.Id(),
	// })
	subscribers := s.getSubscribers(e.Name(), e.Id())
	for _, subscriber := range subscribers {
		asyncer.AsyncHandle(e, subscriber)
	}
}
