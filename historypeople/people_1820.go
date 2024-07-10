package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1820] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1820][31820] = People_31820
	HistoryPeopleMap[1820][71820] = People_71820
	HistoryPeopleMap[1820][81820] = People_81820
	HistoryPeopleMap[1820][191820] = People_191820
}
