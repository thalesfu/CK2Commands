package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3912] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3912][3912] = People_3912
	HistoryPeopleMap[3912][33912] = People_33912
	HistoryPeopleMap[3912][73912] = People_73912
	HistoryPeopleMap[3912][183912] = People_183912
}
