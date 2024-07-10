package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8500] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8500][8500] = People_8500
	HistoryPeopleMap[8500][108500] = People_108500
	HistoryPeopleMap[8500][138500] = People_138500
	HistoryPeopleMap[8500][168500] = People_168500
	HistoryPeopleMap[8500][188500] = People_188500
	HistoryPeopleMap[8500][208500] = People_208500
	HistoryPeopleMap[8500][218500] = People_218500
	HistoryPeopleMap[8500][468500] = People_468500
}
