package actionstats

import "sync"

// NewActionStore initializes an empty action store.
func NewActionStore() *ActionStore {
	return &ActionStore{
		Store: map[string]Stats{},
	}
}

// ActionStore collects stats per action in memory.
type ActionStore struct {
	sync.RWMutex
	Store map[string]Stats
}

type Stats struct {
	Avg   int
	Count int
	Sum   int
}

// GetActionStats returns a map dereference pair (value at key, existence of key) for a given action.
func (a *ActionStore) GetActionStats(action string) (value Stats, ok bool) {
	a.RLock()
	val, ok := a.Store[action]
	a.RUnlock()
	return val, ok
}

func (a *ActionStore) GetAllActionStats() map[string]Stats {
	return a.Store
}

// MergeNewAction performs concurrency-safe update/initialization to stats map value
func (a *ActionStore) MergeNewAction(action string, time int) {
	stats, ok := a.GetActionStats(action)
	a.Lock()
	if !ok {
		a.Store[action] = Stats{
			Avg:   time,
			Count: 1,
			Sum:   time,
		}
	} else {
		stats.Count += 1
		stats.Sum = stats.Sum + time
		stats.Avg = stats.Sum / stats.Count
		a.Store[action] = stats
	}
	a.Unlock()
}
