package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9400] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9400][159400] = People_159400
	HistoryPeopleMap[9400][189400] = People_189400
}
