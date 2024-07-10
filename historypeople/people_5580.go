package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5580] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5580][205580] = People_205580
	HistoryPeopleMap[5580][455580] = People_455580
}
