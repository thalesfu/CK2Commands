package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1887] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1887][31887] = People_31887
	HistoryPeopleMap[1887][71887] = People_71887
	HistoryPeopleMap[1887][191887] = People_191887
}
