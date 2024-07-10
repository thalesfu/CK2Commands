package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4488] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4488][34488] = People_34488
	HistoryPeopleMap[4488][74488] = People_74488
}
