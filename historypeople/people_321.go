package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[321] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[321][321] = People_321
	HistoryPeopleMap[321][20321] = People_20321
	HistoryPeopleMap[321][30321] = People_30321
	HistoryPeopleMap[321][170321] = People_170321
	HistoryPeopleMap[321][190321] = People_190321
	HistoryPeopleMap[321][200321] = People_200321
	HistoryPeopleMap[321][260321] = People_260321
}
