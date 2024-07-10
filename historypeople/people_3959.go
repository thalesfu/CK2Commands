package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3959] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3959][73959] = People_73959
}
