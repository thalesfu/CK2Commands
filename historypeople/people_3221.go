package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3221] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3221][3221] = People_3221
	HistoryPeopleMap[3221][33221] = People_33221
	HistoryPeopleMap[3221][73221] = People_73221
	HistoryPeopleMap[3221][83221] = People_83221
	HistoryPeopleMap[3221][93221] = People_93221
	HistoryPeopleMap[3221][183221] = People_183221
}
