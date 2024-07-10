package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[443] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[443][30443] = People_30443
	HistoryPeopleMap[443][40443] = People_40443
	HistoryPeopleMap[443][190443] = People_190443
	HistoryPeopleMap[443][260443] = People_260443
}
