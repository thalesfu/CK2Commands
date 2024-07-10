package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5625] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5625][145625] = People_145625
	HistoryPeopleMap[5625][205625] = People_205625
}
