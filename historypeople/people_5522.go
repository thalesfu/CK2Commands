package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5522] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5522][125522] = People_125522
	HistoryPeopleMap[5522][205522] = People_205522
	HistoryPeopleMap[5522][455522] = People_455522
	HistoryPeopleMap[5522][465522] = People_465522
}
