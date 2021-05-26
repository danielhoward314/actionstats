package actionstats

import (
	"encoding/json"
	"reflect"
	"sort"
	"sync"
	"testing"
)

func fromStatsJSON(t testing.TB, s string) (stats []ActionAvg, e error) {
	t.Helper()
	b := []byte(s)
	var statsMsg []ActionAvg
	err := json.Unmarshal(b, &statsMsg)
	return statsMsg, err
}

func assertStatsEquivalence(t testing.TB, got, want []ActionAvg) {
	t.Helper()
	if (got == nil) != (want == nil) {
		t.Errorf("Of got %v and want %v, one is nil while the other is not", got, want)
	}
	if len(got) != len(want) {
		t.Errorf("Slice length mismatch: got has %v and want has %v elements", len(got), len(want))
	}
	// reflect.DeepEqual assumes slices are sorted
	sort.Slice(got, func(p, q int) bool {
		return got[p].Action < got[q].Action
	})
	sort.Slice(want, func(p, q int) bool {
		return want[p].Action < want[q].Action
	})
	eq := reflect.DeepEqual(got, want)
	if !eq {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestGetStats(t *testing.T) {
	store := NewActionStore()
	t.Run("it returns array of all averages by action", func(t *testing.T) {
		runTestTable(t, store)
		gotJson := store.GetStats()
		// unmarshaling for testing correct values within the json
		// rather than testing strict equivalence of two json strings
		got, err := fromStatsJSON(t, gotJson)
		want := tt.out
		if err != nil {
			t.Errorf("Got non-nil error %v but want %v", err, want)
		}
		assertStatsEquivalence(t, got, want)
	})
}

func TestGetStatsConcurrently(t *testing.T) {
	store := NewActionStore()
	t.Run("it should allow concurrent calls", func(t *testing.T) {
		var wg sync.WaitGroup
		for _, s := range tt.in {
			wg.Add(1)
			current := s
			go func() {
				defer wg.Done()
				store.AddAction(current)
				// e := store.AddAction(current)
				// if e != nil {
				// 	t.Fatal("should add each action to store without error:", e)
				// }
			}()
		}
		wg.Wait()
		gotJson := store.GetStats()
		got, err := fromStatsJSON(t, gotJson)
		want := tt.out
		if err != nil {
			t.Errorf("Got non-nil error %v but want %v", err, want)
		}
		assertStatsEquivalence(t, got, want)
	})
}
