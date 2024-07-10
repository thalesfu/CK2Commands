package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1612] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1612][1612] = People_1612
	HistoryPeopleMap[1612][71612] = People_71612
	HistoryPeopleMap[1612][161612] = People_161612
	HistoryPeopleMap[1612][191612] = People_191612
}
