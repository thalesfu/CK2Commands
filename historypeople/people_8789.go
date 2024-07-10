package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8789] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8789][168789] = People_168789
	HistoryPeopleMap[8789][188789] = People_188789
	HistoryPeopleMap[8789][248789] = People_248789
}
