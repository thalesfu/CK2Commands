package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1620] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1620][1620] = People_1620
	HistoryPeopleMap[1620][71620] = People_71620
	HistoryPeopleMap[1620][161620] = People_161620
	HistoryPeopleMap[1620][191620] = People_191620
}
