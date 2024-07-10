package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1840] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1840][31840] = People_31840
	HistoryPeopleMap[1840][71840] = People_71840
	HistoryPeopleMap[1840][191840] = People_191840
}
