package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[201] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[201][30201] = People_30201
	HistoryPeopleMap[201][70201] = People_70201
	HistoryPeopleMap[201][160201] = People_160201
	HistoryPeopleMap[201][170201] = People_170201
	HistoryPeopleMap[201][180201] = People_180201
	HistoryPeopleMap[201][190201] = People_190201
	HistoryPeopleMap[201][200201] = People_200201
	HistoryPeopleMap[201][470201] = People_470201
}
