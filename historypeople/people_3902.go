package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3902] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3902][3902] = People_3902
	HistoryPeopleMap[3902][33902] = People_33902
	HistoryPeopleMap[3902][73902] = People_73902
	HistoryPeopleMap[3902][183902] = People_183902
}
