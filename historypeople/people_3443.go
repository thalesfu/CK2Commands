package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3443] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3443][73443] = People_73443
	HistoryPeopleMap[3443][93443] = People_93443
}
