package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[240] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[240][240] = People_240
	HistoryPeopleMap[240][20240] = People_20240
	HistoryPeopleMap[240][30240] = People_30240
	HistoryPeopleMap[240][40240] = People_40240
	HistoryPeopleMap[240][70240] = People_70240
	HistoryPeopleMap[240][170240] = People_170240
	HistoryPeopleMap[240][180240] = People_180240
	HistoryPeopleMap[240][190240] = People_190240
	HistoryPeopleMap[240][200240] = People_200240
	HistoryPeopleMap[240][260240] = People_260240
}
