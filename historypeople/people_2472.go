package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2472] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2472][72472] = People_72472
	HistoryPeopleMap[2472][142472] = People_142472
}
