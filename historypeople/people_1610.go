package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1610] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1610][1610] = People_1610
	HistoryPeopleMap[1610][71610] = People_71610
	HistoryPeopleMap[1610][161610] = People_161610
	HistoryPeopleMap[1610][191610] = People_191610
}
