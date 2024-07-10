package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2646] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2646][32646] = People_32646
	HistoryPeopleMap[2646][72646] = People_72646
	HistoryPeopleMap[2646][142646] = People_142646
}
