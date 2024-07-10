package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1600] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1600][41600] = People_41600
	HistoryPeopleMap[1600][71600] = People_71600
	HistoryPeopleMap[1600][161600] = People_161600
	HistoryPeopleMap[1600][191600] = People_191600
	HistoryPeopleMap[1600][451600] = People_451600
	HistoryPeopleMap[1600][471600] = People_471600
}
