package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1996] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1996][71996] = People_71996
	HistoryPeopleMap[1996][191996] = People_191996
	HistoryPeopleMap[1996][451996] = People_451996
}
