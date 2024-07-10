package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1666] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1666][31666] = People_31666
	HistoryPeopleMap[1666][71666] = People_71666
	HistoryPeopleMap[1666][161666] = People_161666
	HistoryPeopleMap[1666][191666] = People_191666
}
