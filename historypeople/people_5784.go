package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5784] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5784][205784] = People_205784
	HistoryPeopleMap[5784][455784] = People_455784
}
