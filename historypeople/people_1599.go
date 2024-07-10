package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1599] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1599][71599] = People_71599
	HistoryPeopleMap[1599][191599] = People_191599
}
