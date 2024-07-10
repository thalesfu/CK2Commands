package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[221] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[221][221] = People_221
	HistoryPeopleMap[221][30221] = People_30221
	HistoryPeopleMap[221][70221] = People_70221
	HistoryPeopleMap[221][170221] = People_170221
	HistoryPeopleMap[221][190221] = People_190221
	HistoryPeopleMap[221][200221] = People_200221
	HistoryPeopleMap[221][260221] = People_260221
}
