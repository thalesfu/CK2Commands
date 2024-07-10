package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1839] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1839][31839] = People_31839
	HistoryPeopleMap[1839][191839] = People_191839
}
