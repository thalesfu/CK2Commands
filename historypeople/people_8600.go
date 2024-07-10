package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8600] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8600][8600] = People_8600
	HistoryPeopleMap[8600][108600] = People_108600
	HistoryPeopleMap[8600][168600] = People_168600
	HistoryPeopleMap[8600][188600] = People_188600
}
