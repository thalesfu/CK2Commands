package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5250] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5250][125250] = People_125250
	HistoryPeopleMap[5250][145250] = People_145250
	HistoryPeopleMap[5250][155250] = People_155250
	HistoryPeopleMap[5250][205250] = People_205250
}
