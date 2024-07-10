package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[945] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[945][945] = People_945
	HistoryPeopleMap[945][70945] = People_70945
	HistoryPeopleMap[945][260945] = People_260945
}
