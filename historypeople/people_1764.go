package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1764] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1764][1764] = People_1764
	HistoryPeopleMap[1764][31764] = People_31764
	HistoryPeopleMap[1764][71764] = People_71764
	HistoryPeopleMap[1764][161764] = People_161764
	HistoryPeopleMap[1764][191764] = People_191764
}
