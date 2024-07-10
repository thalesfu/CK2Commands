package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1760] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1760][1760] = People_1760
	HistoryPeopleMap[1760][31760] = People_31760
	HistoryPeopleMap[1760][71760] = People_71760
	HistoryPeopleMap[1760][161760] = People_161760
	HistoryPeopleMap[1760][191760] = People_191760
	HistoryPeopleMap[1760][221760] = People_221760
}
