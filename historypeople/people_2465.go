package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2465] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2465][72465] = People_72465
	HistoryPeopleMap[2465][142465] = People_142465
}
