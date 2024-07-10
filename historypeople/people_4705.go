package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4705] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4705][34705] = People_34705
	HistoryPeopleMap[4705][74705] = People_74705
	HistoryPeopleMap[4705][144705] = People_144705
}
