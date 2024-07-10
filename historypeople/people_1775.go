package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1775] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1775][31775] = People_31775
	HistoryPeopleMap[1775][71775] = People_71775
	HistoryPeopleMap[1775][161775] = People_161775
}
