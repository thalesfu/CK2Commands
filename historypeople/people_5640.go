package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5640] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5640][5640] = People_5640
	HistoryPeopleMap[5640][145640] = People_145640
	HistoryPeopleMap[5640][205640] = People_205640
	HistoryPeopleMap[5640][455640] = People_455640
}
