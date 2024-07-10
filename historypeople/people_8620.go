package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8620] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8620][8620] = People_8620
	HistoryPeopleMap[8620][168620] = People_168620
	HistoryPeopleMap[8620][188620] = People_188620
	HistoryPeopleMap[8620][248620] = People_248620
}
