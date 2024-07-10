package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[925] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[925][925] = People_925
	HistoryPeopleMap[925][70925] = People_70925
	HistoryPeopleMap[925][260925] = People_260925
}
