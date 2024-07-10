package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1893] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1893][31893] = People_31893
	HistoryPeopleMap[1893][71893] = People_71893
	HistoryPeopleMap[1893][191893] = People_191893
}
