package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5997] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5997][5997] = People_5997
	HistoryPeopleMap[5997][125997] = People_125997
	HistoryPeopleMap[5997][145997] = People_145997
}
