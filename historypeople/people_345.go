package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[345] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[345][345] = People_345
	HistoryPeopleMap[345][30345] = People_30345
	HistoryPeopleMap[345][40345] = People_40345
	HistoryPeopleMap[345][170345] = People_170345
	HistoryPeopleMap[345][190345] = People_190345
	HistoryPeopleMap[345][200345] = People_200345
	HistoryPeopleMap[345][260345] = People_260345
}
