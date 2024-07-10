package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1780] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1780][1780] = People_1780
	HistoryPeopleMap[1780][31780] = People_31780
	HistoryPeopleMap[1780][191780] = People_191780
}
