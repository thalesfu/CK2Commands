package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1753] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1753][31753] = People_31753
	HistoryPeopleMap[1753][71753] = People_71753
	HistoryPeopleMap[1753][161753] = People_161753
	HistoryPeopleMap[1753][191753] = People_191753
}
