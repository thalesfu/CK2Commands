package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1790] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1790][1790] = People_1790
	HistoryPeopleMap[1790][31790] = People_31790
	HistoryPeopleMap[1790][191790] = People_191790
}
