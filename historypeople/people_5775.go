package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5775] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5775][145775] = People_145775
	HistoryPeopleMap[5775][455775] = People_455775
}
