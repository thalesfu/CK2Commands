package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1886] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1886][31886] = People_31886
	HistoryPeopleMap[1886][71886] = People_71886
	HistoryPeopleMap[1886][191886] = People_191886
}
