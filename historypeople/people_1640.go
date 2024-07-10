package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1640] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1640][71640] = People_71640
	HistoryPeopleMap[1640][161640] = People_161640
	HistoryPeopleMap[1640][191640] = People_191640
}
