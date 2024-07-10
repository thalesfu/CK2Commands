package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5720] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5720][125720] = People_125720
	HistoryPeopleMap[5720][145720] = People_145720
	HistoryPeopleMap[5720][455720] = People_455720
}
