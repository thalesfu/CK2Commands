package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3822] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3822][3822] = People_3822
	HistoryPeopleMap[3822][73822] = People_73822
	HistoryPeopleMap[3822][93822] = People_93822
}
