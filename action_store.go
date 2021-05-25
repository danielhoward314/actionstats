package main

import "sync"

// NewActionStore initializes an empty action store.
func NewActionStore() *ActionStore {
	return &ActionStore{
		store: map[string]Stats{},
	}
}

// ActionStore collects stats per action in memory.
type ActionStore struct {
	sync.RWMutex
	store map[string]Stats
}

type Stats struct {
	avg   int
	count int
	sum   int
}

// GetActionStats returns a map dereference pair (value at key, existence of value at key) for a given action.
func (a *ActionStore) GetActionStats(action string) (value Stats, ok bool) {
	a.Lock()
	val, ok := a.store[action]
	a.Unlock()
	return val, ok
}

func (a *ActionStore) GetAllActionStats() map[string]Stats {
	return a.store
}

// performs concurrency-safe update/initialization to stats map value
func (a *ActionStore) MergeNewAction(action string, time int) {
	stats, ok := a.GetActionStats(action)
	a.Lock()
	if !ok {
		a.store[action] = Stats{
			avg:   time,
			count: 1,
			sum:   time,
		}
	} else {
		stats.count += 1
		stats.sum = stats.sum + time
		stats.avg = stats.sum / stats.count
		a.store[action] = stats
	}
	a.Unlock()
}

func (a *ActionStore) AddAction(s string) error {
	aMsg, e := FromJSON(s)
	if e == nil {
		a.MergeNewAction(aMsg.Action, aMsg.Time)
	}
	return e
}

func (a *ActionStore) GetStats() string {
	stats := a.GetAllActionStats()
	var avgs []map[string]int
	for k, v := range stats {
		el := make(map[string]int)
		el[k] = v.avg
		avgs = append(avgs, el)
	}
	return ToJSON(avgs)
}
