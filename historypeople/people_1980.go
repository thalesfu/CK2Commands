package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1980] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1980][71980] = People_71980
	HistoryPeopleMap[1980][161980] = People_161980
	HistoryPeopleMap[1980][191980] = People_191980
	HistoryPeopleMap[1980][451980] = People_451980
}
