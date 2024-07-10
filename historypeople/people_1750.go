package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1750] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1750][1750] = People_1750
	HistoryPeopleMap[1750][31750] = People_31750
	HistoryPeopleMap[1750][71750] = People_71750
	HistoryPeopleMap[1750][161750] = People_161750
	HistoryPeopleMap[1750][191750] = People_191750
	HistoryPeopleMap[1750][221750] = People_221750
}
