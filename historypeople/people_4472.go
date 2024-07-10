package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4472] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4472][34472] = People_34472
	HistoryPeopleMap[4472][74472] = People_74472
}
