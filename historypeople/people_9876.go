package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9876] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9876][159876] = People_159876
}
