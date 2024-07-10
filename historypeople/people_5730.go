package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5730] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5730][125730] = People_125730
	HistoryPeopleMap[5730][145730] = People_145730
	HistoryPeopleMap[5730][205730] = People_205730
	HistoryPeopleMap[5730][455730] = People_455730
}
