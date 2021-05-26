package actionstats

func (a *ActionStore) AddAction(s string) error {
	aMsg, e := FromActionJSON(s)
	if e == nil {
		a.MergeNewAction(aMsg.Action, aMsg.Time)
	}
	return e
}
