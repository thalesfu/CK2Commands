package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9812] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9812][9812] = People_9812
	HistoryPeopleMap[9812][159812] = People_159812
}
