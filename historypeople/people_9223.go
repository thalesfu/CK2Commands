package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9223] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9223][159223] = People_159223
	HistoryPeopleMap[9223][189223] = People_189223
}
