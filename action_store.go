package actionstats

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

// GetActionStats returns a map dereference pair (value at key, existence of key) for a given action.
func (a *ActionStore) GetActionStats(action string) (value Stats, ok bool) {
	a.RLock()
	val, ok := a.store[action]
	a.RUnlock()
	return val, ok
}

func (a *ActionStore) GetAllActionStats() map[string]Stats {
	return a.store
}

// MergeNewAction performs concurrency-safe update/initialization to stats map value
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
