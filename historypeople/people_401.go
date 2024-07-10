package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[401] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[401][401] = People_401
	HistoryPeopleMap[401][30401] = People_30401
	HistoryPeopleMap[401][70401] = People_70401
	HistoryPeopleMap[401][190401] = People_190401
	HistoryPeopleMap[401][260401] = People_260401
}
