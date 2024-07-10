package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3801] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3801][73801] = People_73801
	HistoryPeopleMap[3801][93801] = People_93801
}
