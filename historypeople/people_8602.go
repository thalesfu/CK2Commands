package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8602] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8602][8602] = People_8602
	HistoryPeopleMap[8602][168602] = People_168602
	HistoryPeopleMap[8602][188602] = People_188602
	HistoryPeopleMap[8602][248602] = People_248602
}
