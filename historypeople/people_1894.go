package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1894] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1894][71894] = People_71894
	HistoryPeopleMap[1894][191894] = People_191894
}
