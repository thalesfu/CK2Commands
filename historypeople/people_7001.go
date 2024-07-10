package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7001] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7001][127001] = People_127001
	HistoryPeopleMap[7001][147001] = People_147001
	HistoryPeopleMap[7001][167001] = People_167001
	HistoryPeopleMap[7001][187001] = People_187001
	HistoryPeopleMap[7001][227001] = People_227001
	HistoryPeopleMap[7001][247001] = People_247001
	HistoryPeopleMap[7001][487001] = People_487001
}
