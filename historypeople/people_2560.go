package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2560] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2560][32560] = People_32560
	HistoryPeopleMap[2560][72560] = People_72560
	HistoryPeopleMap[2560][102560] = People_102560
	HistoryPeopleMap[2560][142560] = People_142560
	HistoryPeopleMap[2560][212560] = People_212560
	HistoryPeopleMap[2560][222560] = People_222560
	HistoryPeopleMap[2560][232560] = People_232560
}
