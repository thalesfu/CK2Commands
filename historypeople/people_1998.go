package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1998] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1998][71998] = People_71998
	HistoryPeopleMap[1998][191998] = People_191998
	HistoryPeopleMap[1998][451998] = People_451998
}
