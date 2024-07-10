package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1784] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1784][31784] = People_31784
	HistoryPeopleMap[1784][191784] = People_191784
}
