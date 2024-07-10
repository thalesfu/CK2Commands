package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5465] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5465][455465] = People_455465
}
