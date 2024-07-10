package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5865] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5865][145865] = People_145865
}
