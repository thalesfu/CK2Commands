package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1852] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1852][31852] = People_31852
	HistoryPeopleMap[1852][71852] = People_71852
	HistoryPeopleMap[1852][191852] = People_191852
}
