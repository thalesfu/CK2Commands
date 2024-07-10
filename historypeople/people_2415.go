package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2415] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2415][12415] = People_12415
	HistoryPeopleMap[2415][72415] = People_72415
	HistoryPeopleMap[2415][142415] = People_142415
}
