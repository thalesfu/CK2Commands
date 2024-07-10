package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[320] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[320][320] = People_320
	HistoryPeopleMap[320][20320] = People_20320
	HistoryPeopleMap[320][30320] = People_30320
	HistoryPeopleMap[320][40320] = People_40320
	HistoryPeopleMap[320][170320] = People_170320
	HistoryPeopleMap[320][190320] = People_190320
	HistoryPeopleMap[320][200320] = People_200320
	HistoryPeopleMap[320][260320] = People_260320
}
