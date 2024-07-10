package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4223] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4223][4223] = People_4223
	HistoryPeopleMap[4223][34223] = People_34223
	HistoryPeopleMap[4223][144223] = People_144223
	HistoryPeopleMap[4223][194223] = People_194223
}
