package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1896] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1896][191896] = People_191896
}
