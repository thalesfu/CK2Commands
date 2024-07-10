package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3805] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3805][73805] = People_73805
	HistoryPeopleMap[3805][93805] = People_93805
}
