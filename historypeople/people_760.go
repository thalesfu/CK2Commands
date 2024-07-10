package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[760] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[760][760] = People_760
	HistoryPeopleMap[760][30760] = People_30760
	HistoryPeopleMap[760][70760] = People_70760
	HistoryPeopleMap[760][180760] = People_180760
	HistoryPeopleMap[760][260760] = People_260760
	HistoryPeopleMap[760][450760] = People_450760
}
