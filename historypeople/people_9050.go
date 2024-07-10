package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9050] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9050][9050] = People_9050
	HistoryPeopleMap[9050][129050] = People_129050
	HistoryPeopleMap[9050][159050] = People_159050
	HistoryPeopleMap[9050][189050] = People_189050
	HistoryPeopleMap[9050][479050] = People_479050
}
