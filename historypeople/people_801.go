package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[801] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[801][801] = People_801
	HistoryPeopleMap[801][20801] = People_20801
	HistoryPeopleMap[801][30801] = People_30801
	HistoryPeopleMap[801][40801] = People_40801
	HistoryPeopleMap[801][70801] = People_70801
	HistoryPeopleMap[801][220801] = People_220801
	HistoryPeopleMap[801][250801] = People_250801
	HistoryPeopleMap[801][450801] = People_450801
}
