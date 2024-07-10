package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4160] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4160][144160] = People_144160
	HistoryPeopleMap[4160][194160] = People_194160
	HistoryPeopleMap[4160][454160] = People_454160
}
