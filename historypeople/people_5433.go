package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5433] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5433][455433] = People_455433
}
