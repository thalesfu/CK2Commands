package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[234] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[234][234] = People_234
	HistoryPeopleMap[234][20234] = People_20234
	HistoryPeopleMap[234][30234] = People_30234
	HistoryPeopleMap[234][40234] = People_40234
	HistoryPeopleMap[234][70234] = People_70234
	HistoryPeopleMap[234][170234] = People_170234
	HistoryPeopleMap[234][180234] = People_180234
	HistoryPeopleMap[234][190234] = People_190234
	HistoryPeopleMap[234][200234] = People_200234
	HistoryPeopleMap[234][260234] = People_260234
}
