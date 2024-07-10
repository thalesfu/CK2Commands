package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3600] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3600][33600] = People_33600
	HistoryPeopleMap[3600][73600] = People_73600
	HistoryPeopleMap[3600][83600] = People_83600
	HistoryPeopleMap[3600][213600] = People_213600
	HistoryPeopleMap[3600][223600] = People_223600
	HistoryPeopleMap[3600][453600] = People_453600
	HistoryPeopleMap[3600][463600] = People_463600
}
