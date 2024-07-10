package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[512] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[512][512] = People_512
	HistoryPeopleMap[512][20512] = People_20512
	HistoryPeopleMap[512][30512] = People_30512
	HistoryPeopleMap[512][40512] = People_40512
	HistoryPeopleMap[512][70512] = People_70512
	HistoryPeopleMap[512][100512] = People_100512
	HistoryPeopleMap[512][110512] = People_110512
	HistoryPeopleMap[512][180512] = People_180512
	HistoryPeopleMap[512][190512] = People_190512
	HistoryPeopleMap[512][210512] = People_210512
	HistoryPeopleMap[512][260512] = People_260512
	HistoryPeopleMap[512][450512] = People_450512
}
