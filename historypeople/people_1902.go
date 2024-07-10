package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1902] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1902][71902] = People_71902
	HistoryPeopleMap[1902][161902] = People_161902
	HistoryPeopleMap[1902][191902] = People_191902
	HistoryPeopleMap[1902][461902] = People_461902
}
