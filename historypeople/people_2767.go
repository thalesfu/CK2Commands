package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2767] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2767][32767] = People_32767
	HistoryPeopleMap[2767][72767] = People_72767
	HistoryPeopleMap[2767][142767] = People_142767
	HistoryPeopleMap[2767][232767] = People_232767
}
