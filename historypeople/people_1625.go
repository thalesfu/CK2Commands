package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1625] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1625][71625] = People_71625
	HistoryPeopleMap[1625][161625] = People_161625
	HistoryPeopleMap[1625][191625] = People_191625
}
