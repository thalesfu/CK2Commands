package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9600] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9600][109600] = People_109600
	HistoryPeopleMap[9600][159600] = People_159600
}
