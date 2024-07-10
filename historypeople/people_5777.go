package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5777] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5777][145777] = People_145777
	HistoryPeopleMap[5777][455777] = People_455777
}
