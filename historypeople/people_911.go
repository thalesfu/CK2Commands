package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[911] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[911][911] = People_911
	HistoryPeopleMap[911][30911] = People_30911
	HistoryPeopleMap[911][70911] = People_70911
	HistoryPeopleMap[911][260911] = People_260911
}
