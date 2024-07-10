package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1748] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1748][1748] = People_1748
	HistoryPeopleMap[1748][71748] = People_71748
	HistoryPeopleMap[1748][161748] = People_161748
	HistoryPeopleMap[1748][191748] = People_191748
}
