package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5611] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5611][145611] = People_145611
	HistoryPeopleMap[5611][455611] = People_455611
}
