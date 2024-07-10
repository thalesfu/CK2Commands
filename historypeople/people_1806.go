package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1806] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1806][31806] = People_31806
	HistoryPeopleMap[1806][71806] = People_71806
	HistoryPeopleMap[1806][191806] = People_191806
}
