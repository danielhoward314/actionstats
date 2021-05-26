package actionstats

import (
	"reflect"
	"testing"
)

var store *ActionStore

type testCase struct {
	in     []string
	effect map[string]Stats
	out    []ActionAvg
}

var tt = testCase{
	in: []string{
		`{"action":"jump", "time":100}`,
		`{"action":"dance", "time":122}`,
		`{"action":"run", "time":75}`,
		`{"action":"jump", "time":200}`,
		`{"action":"run", "time":85}`,
		`{"action":"jump", "time":300}`,
		`{"action":"dance", "time":222}`,
		`{"action":"jump", "time":400}`,
		`{"action":"run", "time":95}`,
		`{"action":"run", "time":105}`,
		`{"action":"sing", "time":126}`,
		`{"action":"jump", "time":500}`,
	},
	effect: map[string]Stats{
		"dance": {Avg: 172, Count: 2, Sum: 344},
		"jump":  {Avg: 300, Count: 5, Sum: 1500},
		"run":   {Avg: 90, Count: 4, Sum: 360},
		"sing":  {Avg: 126, Count: 1, Sum: 126},
	},
	out: []ActionAvg{
		{Action: "dance", Avg: 172},
		{Action: "jump", Avg: 300},
		{Action: "run", Avg: 90},
		{Action: "sing", Avg: 126},
	},
}

func runTestTable(t testing.TB, store *ActionStore) {
	t.Helper()
	for _, s := range tt.in {
		e := store.AddAction(s)
		if e != nil {
			t.Fatal("should add each action to store without error:", e)
		}
	}
}

func assertStoreEquivalence(t testing.TB, got, want map[string]Stats) {
	t.Helper()
	eq := reflect.DeepEqual(got, want)
	if !eq {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestAddActionSuccess(t *testing.T) {
	store = NewActionStore()
	t.Run("it stores by action the stats needed to track avg", func(t *testing.T) {
		runTestTable(t, store)
		// assert on desired side-effect in memory
		assertStoreEquivalence(t, store.GetAllActionStats(), tt.effect)
	})
}

func TestAddActionError(t *testing.T) {
	store = NewActionStore()
	malformed := `{action":jump" time":}}100}`
	t.Run("it returns an error", func(t *testing.T) {
		e := store.AddAction(malformed)
		if e == nil {
			t.Fatal("should return non-nil error for malformed json:", e)
		}
		switch e.(type) {
		case error:
			return
		default:
			t.Fatal("should have return value with type error")
		}
	})
}
