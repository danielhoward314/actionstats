package actionstats

func (a *ActionStore) GetStats() string {
	stats := a.GetAllActionStats()
	var avgs []ActionAvg
	for k, v := range stats {
		el := ActionAvg{
			Action: k,
			Avg:    v.Avg,
		}
		avgs = append(avgs, el)
	}
	return ToStatsJSON(avgs)
}
