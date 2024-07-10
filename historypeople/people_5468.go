package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5468] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5468][455468] = People_455468
}
