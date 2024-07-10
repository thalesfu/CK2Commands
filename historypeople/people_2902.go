package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2902] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2902][32902] = People_32902
	HistoryPeopleMap[2902][72902] = People_72902
	HistoryPeopleMap[2902][142902] = People_142902
}
