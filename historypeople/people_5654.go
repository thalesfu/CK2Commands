package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5654] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5654][145654] = People_145654
	HistoryPeopleMap[5654][455654] = People_455654
}
