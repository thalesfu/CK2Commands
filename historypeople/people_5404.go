package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5404] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5404][125404] = People_125404
	HistoryPeopleMap[5404][455404] = People_455404
}
