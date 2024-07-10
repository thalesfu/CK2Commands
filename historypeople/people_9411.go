package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9411] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9411][159411] = People_159411
	HistoryPeopleMap[9411][189411] = People_189411
}
