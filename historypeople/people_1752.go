package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1752] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1752][1752] = People_1752
	HistoryPeopleMap[1752][31752] = People_31752
	HistoryPeopleMap[1752][71752] = People_71752
	HistoryPeopleMap[1752][161752] = People_161752
	HistoryPeopleMap[1752][191752] = People_191752
}
