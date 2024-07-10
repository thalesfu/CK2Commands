package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8250] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8250][138250] = People_138250
	HistoryPeopleMap[8250][168250] = People_168250
	HistoryPeopleMap[8250][188250] = People_188250
}
