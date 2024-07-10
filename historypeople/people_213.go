package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[213] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[213][213] = People_213
	HistoryPeopleMap[213][20213] = People_20213
	HistoryPeopleMap[213][30213] = People_30213
	HistoryPeopleMap[213][170213] = People_170213
	HistoryPeopleMap[213][180213] = People_180213
	HistoryPeopleMap[213][190213] = People_190213
	HistoryPeopleMap[213][200213] = People_200213
	HistoryPeopleMap[213][470213] = People_470213
}
