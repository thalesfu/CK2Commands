package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9300] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9300][9300] = People_9300
	HistoryPeopleMap[9300][159300] = People_159300
}
