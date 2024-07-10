package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9781] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9781][159781] = People_159781
}
