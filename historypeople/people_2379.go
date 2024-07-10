package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2379] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2379][12379] = People_12379
	HistoryPeopleMap[2379][72379] = People_72379
	HistoryPeopleMap[2379][142379] = People_142379
	HistoryPeopleMap[2379][252379] = People_252379
}
