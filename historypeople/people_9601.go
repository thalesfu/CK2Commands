package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9601] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9601][109601] = People_109601
	HistoryPeopleMap[9601][159601] = People_159601
}
