package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8760] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8760][168760] = People_168760
	HistoryPeopleMap[8760][188760] = People_188760
	HistoryPeopleMap[8760][248760] = People_248760
}
