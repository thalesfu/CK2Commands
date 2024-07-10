package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[940] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[940][940] = People_940
	HistoryPeopleMap[940][70940] = People_70940
	HistoryPeopleMap[940][260940] = People_260940
}
