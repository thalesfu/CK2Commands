package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5630] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5630][145630] = People_145630
	HistoryPeopleMap[5630][205630] = People_205630
	HistoryPeopleMap[5630][455630] = People_455630
}
