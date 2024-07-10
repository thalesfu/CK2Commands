package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5620] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5620][145620] = People_145620
	HistoryPeopleMap[5620][205620] = People_205620
	HistoryPeopleMap[5620][455620] = People_455620
}
