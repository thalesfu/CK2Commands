package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2600] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2600][32600] = People_32600
	HistoryPeopleMap[2600][72600] = People_72600
	HistoryPeopleMap[2600][112600] = People_112600
	HistoryPeopleMap[2600][142600] = People_142600
	HistoryPeopleMap[2600][212600] = People_212600
	HistoryPeopleMap[2600][222600] = People_222600
	HistoryPeopleMap[2600][232600] = People_232600
	HistoryPeopleMap[2600][452600] = People_452600
	HistoryPeopleMap[2600][462600] = People_462600
}
