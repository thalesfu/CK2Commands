package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1870] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1870][31870] = People_31870
	HistoryPeopleMap[1870][71870] = People_71870
	HistoryPeopleMap[1870][191870] = People_191870
}
