package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5403] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5403][125403] = People_125403
	HistoryPeopleMap[5403][455403] = People_455403
}
