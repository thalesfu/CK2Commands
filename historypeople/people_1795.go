package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1795] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1795][31795] = People_31795
	HistoryPeopleMap[1795][191795] = People_191795
}
