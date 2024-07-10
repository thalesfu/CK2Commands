package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1773] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1773][1773] = People_1773
	HistoryPeopleMap[1773][31773] = People_31773
	HistoryPeopleMap[1773][71773] = People_71773
	HistoryPeopleMap[1773][161773] = People_161773
}
