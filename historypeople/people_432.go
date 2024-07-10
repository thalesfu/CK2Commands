package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[432] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[432][432] = People_432
	HistoryPeopleMap[432][30432] = People_30432
	HistoryPeopleMap[432][40432] = People_40432
	HistoryPeopleMap[432][70432] = People_70432
	HistoryPeopleMap[432][190432] = People_190432
	HistoryPeopleMap[432][260432] = People_260432
}
