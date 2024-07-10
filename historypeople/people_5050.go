package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5050] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5050][45050] = People_45050
	HistoryPeopleMap[5050][125050] = People_125050
	HistoryPeopleMap[5050][145050] = People_145050
	HistoryPeopleMap[5050][155050] = People_155050
	HistoryPeopleMap[5050][175050] = People_175050
	HistoryPeopleMap[5050][205050] = People_205050
	HistoryPeopleMap[5050][235050] = People_235050
	HistoryPeopleMap[5050][485050] = People_485050
}
