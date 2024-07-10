package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1884] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1884][31884] = People_31884
	HistoryPeopleMap[1884][71884] = People_71884
	HistoryPeopleMap[1884][191884] = People_191884
}
