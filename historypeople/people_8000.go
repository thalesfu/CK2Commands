package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8000] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8000][98000] = People_98000
	HistoryPeopleMap[8000][138000] = People_138000
	HistoryPeopleMap[8000][168000] = People_168000
	HistoryPeopleMap[8000][178000] = People_178000
	HistoryPeopleMap[8000][188000] = People_188000
	HistoryPeopleMap[8000][218000] = People_218000
	HistoryPeopleMap[8000][228000] = People_228000
	HistoryPeopleMap[8000][478000] = People_478000
}
