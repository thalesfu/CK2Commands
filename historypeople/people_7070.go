package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7070] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7070][127070] = People_127070
	HistoryPeopleMap[7070][247070] = People_247070
	HistoryPeopleMap[7070][487070] = People_487070
}
