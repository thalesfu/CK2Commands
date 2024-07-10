package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5789] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5789][5789] = People_5789
	HistoryPeopleMap[5789][145789] = People_145789
	HistoryPeopleMap[5789][205789] = People_205789
	HistoryPeopleMap[5789][455789] = People_455789
}
