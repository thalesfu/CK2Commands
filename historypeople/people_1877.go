package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1877] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1877][31877] = People_31877
	HistoryPeopleMap[1877][71877] = People_71877
	HistoryPeopleMap[1877][191877] = People_191877
}
