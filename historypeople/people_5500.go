package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5500] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5500][125500] = People_125500
	HistoryPeopleMap[5500][205500] = People_205500
	HistoryPeopleMap[5500][215500] = People_215500
	HistoryPeopleMap[5500][455500] = People_455500
	HistoryPeopleMap[5500][465500] = People_465500
}
