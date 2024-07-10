package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[500] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[500][500] = People_500
	HistoryPeopleMap[500][30500] = People_30500
	HistoryPeopleMap[500][40500] = People_40500
	HistoryPeopleMap[500][70500] = People_70500
	HistoryPeopleMap[500][100500] = People_100500
	HistoryPeopleMap[500][110500] = People_110500
	HistoryPeopleMap[500][180500] = People_180500
	HistoryPeopleMap[500][190500] = People_190500
	HistoryPeopleMap[500][200500] = People_200500
	HistoryPeopleMap[500][210500] = People_210500
	HistoryPeopleMap[500][260500] = People_260500
	HistoryPeopleMap[500][450500] = People_450500
	HistoryPeopleMap[500][460500] = People_460500
}
