package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1809] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1809][31809] = People_31809
	HistoryPeopleMap[1809][71809] = People_71809
	HistoryPeopleMap[1809][191809] = People_191809
}
