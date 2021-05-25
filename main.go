package main

import "fmt"

func main() {
	store := NewActionStore()
	fmt.Println(store.GetActionStats("jump"))
	fmt.Println(store.GetActionStats("run"))
	s := `{"action":"jump", "time":100}`
	store.AddAction(s)
	s1 := `{"action":"run", "time":75}`
	store.AddAction(s1)
	s2 := `{"action":"jump", "time":200}`
	store.AddAction(s2)
	fmt.Println(store.GetActionStats("jump"))
	fmt.Println(store.GetActionStats("run"))

	result := store.GetStats()

	fmt.Println(result)
}
