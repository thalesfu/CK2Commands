package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7000] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7000][127000] = People_127000
	HistoryPeopleMap[7000][147000] = People_147000
	HistoryPeopleMap[7000][157000] = People_157000
	HistoryPeopleMap[7000][167000] = People_167000
	HistoryPeopleMap[7000][187000] = People_187000
	HistoryPeopleMap[7000][227000] = People_227000
	HistoryPeopleMap[7000][247000] = People_247000
	HistoryPeopleMap[7000][487000] = People_487000
}
