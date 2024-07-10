package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[333] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[333][333] = People_333
	HistoryPeopleMap[333][30333] = People_30333
	HistoryPeopleMap[333][40333] = People_40333
	HistoryPeopleMap[333][170333] = People_170333
	HistoryPeopleMap[333][200333] = People_200333
	HistoryPeopleMap[333][260333] = People_260333
}
