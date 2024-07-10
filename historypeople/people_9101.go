package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9101] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9101][159101] = People_159101
	HistoryPeopleMap[9101][189101] = People_189101
}
