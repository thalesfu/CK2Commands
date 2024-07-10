package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7003] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7003][7003] = People_7003
	HistoryPeopleMap[7003][127003] = People_127003
	HistoryPeopleMap[7003][147003] = People_147003
	HistoryPeopleMap[7003][167003] = People_167003
	HistoryPeopleMap[7003][187003] = People_187003
	HistoryPeopleMap[7003][227003] = People_227003
	HistoryPeopleMap[7003][247003] = People_247003
}
