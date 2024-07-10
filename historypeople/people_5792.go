package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5792] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5792][5792] = People_5792
	HistoryPeopleMap[5792][145792] = People_145792
	HistoryPeopleMap[5792][455792] = People_455792
}
