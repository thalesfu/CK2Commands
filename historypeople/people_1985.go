package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1985] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1985][71985] = People_71985
	HistoryPeopleMap[1985][161985] = People_161985
	HistoryPeopleMap[1985][191985] = People_191985
	HistoryPeopleMap[1985][451985] = People_451985
}
