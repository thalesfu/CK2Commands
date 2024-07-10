package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1492] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1492][31492] = People_31492
	HistoryPeopleMap[1492][71492] = People_71492
	HistoryPeopleMap[1492][91492] = People_91492
	HistoryPeopleMap[1492][161492] = People_161492
	HistoryPeopleMap[1492][191492] = People_191492
}
