package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8192] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8192][138192] = People_138192
	HistoryPeopleMap[8192][168192] = People_168192
	HistoryPeopleMap[8192][188192] = People_188192
	HistoryPeopleMap[8192][248192] = People_248192
}
