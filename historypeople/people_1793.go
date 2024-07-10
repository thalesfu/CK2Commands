package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1793] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1793][31793] = People_31793
	HistoryPeopleMap[1793][191793] = People_191793
}
