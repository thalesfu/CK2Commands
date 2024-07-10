package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1440] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1440][31440] = People_31440
	HistoryPeopleMap[1440][71440] = People_71440
	HistoryPeopleMap[1440][91440] = People_91440
	HistoryPeopleMap[1440][161440] = People_161440
	HistoryPeopleMap[1440][191440] = People_191440
}
