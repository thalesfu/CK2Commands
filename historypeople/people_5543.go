package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5543] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5543][205543] = People_205543
	HistoryPeopleMap[5543][455543] = People_455543
}
