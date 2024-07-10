package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[645] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[645][70645] = People_70645
	HistoryPeopleMap[645][180645] = People_180645
	HistoryPeopleMap[645][260645] = People_260645
}
