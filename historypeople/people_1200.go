package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1200] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1200][1200] = People_1200
	HistoryPeopleMap[1200][31200] = People_31200
	HistoryPeopleMap[1200][41200] = People_41200
	HistoryPeopleMap[1200][71200] = People_71200
	HistoryPeopleMap[1200][91200] = People_91200
	HistoryPeopleMap[1200][161200] = People_161200
	HistoryPeopleMap[1200][191200] = People_191200
	HistoryPeopleMap[1200][251200] = People_251200
	HistoryPeopleMap[1200][261200] = People_261200
}
