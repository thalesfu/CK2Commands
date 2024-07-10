package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3809] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3809][73809] = People_73809
	HistoryPeopleMap[3809][93809] = People_93809
}
