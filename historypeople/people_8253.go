package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8253] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8253][168253] = People_168253
	HistoryPeopleMap[8253][188253] = People_188253
}
