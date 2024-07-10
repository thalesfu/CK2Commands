package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1878] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1878][31878] = People_31878
	HistoryPeopleMap[1878][71878] = People_71878
	HistoryPeopleMap[1878][191878] = People_191878
}
