package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3803] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3803][73803] = People_73803
	HistoryPeopleMap[3803][93803] = People_93803
}
