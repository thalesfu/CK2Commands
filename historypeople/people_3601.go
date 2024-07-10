package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3601] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3601][33601] = People_33601
	HistoryPeopleMap[3601][73601] = People_73601
	HistoryPeopleMap[3601][83601] = People_83601
	HistoryPeopleMap[3601][213601] = People_213601
	HistoryPeopleMap[3601][223601] = People_223601
	HistoryPeopleMap[3601][453601] = People_453601
}
