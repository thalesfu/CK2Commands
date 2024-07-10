package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[700] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[700][20700] = People_20700
	HistoryPeopleMap[700][30700] = People_30700
	HistoryPeopleMap[700][40700] = People_40700
	HistoryPeopleMap[700][70700] = People_70700
	HistoryPeopleMap[700][260700] = People_260700
	HistoryPeopleMap[700][450700] = People_450700
}
