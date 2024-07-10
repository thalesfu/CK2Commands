package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[980] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[980][980] = People_980
	HistoryPeopleMap[980][70980] = People_70980
	HistoryPeopleMap[980][260980] = People_260980
}
