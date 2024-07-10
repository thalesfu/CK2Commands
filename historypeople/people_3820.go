package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3820] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3820][3820] = People_3820
	HistoryPeopleMap[3820][73820] = People_73820
	HistoryPeopleMap[3820][93820] = People_93820
}
