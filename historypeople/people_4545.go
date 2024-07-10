package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4545] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4545][74545] = People_74545
	HistoryPeopleMap[4545][144545] = People_144545
}
