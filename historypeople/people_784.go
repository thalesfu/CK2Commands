package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[784] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[784][20784] = People_20784
	HistoryPeopleMap[784][30784] = People_30784
	HistoryPeopleMap[784][70784] = People_70784
	HistoryPeopleMap[784][260784] = People_260784
	HistoryPeopleMap[784][450784] = People_450784
}
