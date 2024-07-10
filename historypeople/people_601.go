package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[601] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[601][601] = People_601
	HistoryPeopleMap[601][30601] = People_30601
	HistoryPeopleMap[601][40601] = People_40601
	HistoryPeopleMap[601][70601] = People_70601
	HistoryPeopleMap[601][200601] = People_200601
	HistoryPeopleMap[601][260601] = People_260601
	HistoryPeopleMap[601][450601] = People_450601
}
