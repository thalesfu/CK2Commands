package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1609] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1609][71609] = People_71609
	HistoryPeopleMap[1609][161609] = People_161609
	HistoryPeopleMap[1609][191609] = People_191609
}
