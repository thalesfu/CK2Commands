package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1915] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1915][161915] = People_161915
	HistoryPeopleMap[1915][191915] = People_191915
}
