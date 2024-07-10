package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1994] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1994][71994] = People_71994
	HistoryPeopleMap[1994][191994] = People_191994
	HistoryPeopleMap[1994][451994] = People_451994
}
