package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1875] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1875][31875] = People_31875
	HistoryPeopleMap[1875][71875] = People_71875
	HistoryPeopleMap[1875][191875] = People_191875
}
