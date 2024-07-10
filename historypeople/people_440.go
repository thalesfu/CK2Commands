package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[440] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[440][440] = People_440
	HistoryPeopleMap[440][30440] = People_30440
	HistoryPeopleMap[440][40440] = People_40440
	HistoryPeopleMap[440][70440] = People_70440
	HistoryPeopleMap[440][190440] = People_190440
	HistoryPeopleMap[440][260440] = People_260440
}
