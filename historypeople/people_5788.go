package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5788] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5788][5788] = People_5788
	HistoryPeopleMap[5788][145788] = People_145788
	HistoryPeopleMap[5788][205788] = People_205788
	HistoryPeopleMap[5788][455788] = People_455788
}
