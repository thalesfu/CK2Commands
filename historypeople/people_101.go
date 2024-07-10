package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[101] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[101][101] = People_101
	HistoryPeopleMap[101][30101] = People_30101
	HistoryPeopleMap[101][70101] = People_70101
	HistoryPeopleMap[101][160101] = People_160101
	HistoryPeopleMap[101][170101] = People_170101
	HistoryPeopleMap[101][180101] = People_180101
	HistoryPeopleMap[101][190101] = People_190101
	HistoryPeopleMap[101][200101] = People_200101
	HistoryPeopleMap[101][260101] = People_260101
	HistoryPeopleMap[101][470101] = People_470101
}
