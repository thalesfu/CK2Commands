package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1950] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1950][71950] = People_71950
	HistoryPeopleMap[1950][161950] = People_161950
	HistoryPeopleMap[1950][191950] = People_191950
	HistoryPeopleMap[1950][451950] = People_451950
}
