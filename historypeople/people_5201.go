package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5201] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5201][125201] = People_125201
	HistoryPeopleMap[5201][145201] = People_145201
	HistoryPeopleMap[5201][155201] = People_155201
	HistoryPeopleMap[5201][205201] = People_205201
	HistoryPeopleMap[5201][235201] = People_235201
}
