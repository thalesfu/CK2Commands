package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5769] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5769][145769] = People_145769
	HistoryPeopleMap[5769][455769] = People_455769
}
