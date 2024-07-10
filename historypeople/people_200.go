package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[200] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[200][200] = People_200
	HistoryPeopleMap[200][30200] = People_30200
	HistoryPeopleMap[200][40200] = People_40200
	HistoryPeopleMap[200][70200] = People_70200
	HistoryPeopleMap[200][160200] = People_160200
	HistoryPeopleMap[200][170200] = People_170200
	HistoryPeopleMap[200][180200] = People_180200
	HistoryPeopleMap[200][190200] = People_190200
	HistoryPeopleMap[200][200200] = People_200200
	HistoryPeopleMap[200][470200] = People_470200
	HistoryPeopleMap[200][480200] = People_480200
}
