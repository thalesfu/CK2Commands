package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5261] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5261][145261] = People_145261
	HistoryPeopleMap[5261][205261] = People_205261
}
