package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1835] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1835][31835] = People_31835
	HistoryPeopleMap[1835][71835] = People_71835
	HistoryPeopleMap[1835][191835] = People_191835
}
