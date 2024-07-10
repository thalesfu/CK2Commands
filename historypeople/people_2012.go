package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2012] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2012][2012] = People_2012
	HistoryPeopleMap[2012][12012] = People_12012
	HistoryPeopleMap[2012][42012] = People_42012
	HistoryPeopleMap[2012][72012] = People_72012
	HistoryPeopleMap[2012][82012] = People_82012
	HistoryPeopleMap[2012][142012] = People_142012
	HistoryPeopleMap[2012][182012] = People_182012
	HistoryPeopleMap[2012][192012] = People_192012
	HistoryPeopleMap[2012][202012] = People_202012
	HistoryPeopleMap[2012][212012] = People_212012
	HistoryPeopleMap[2012][252012] = People_252012
	HistoryPeopleMap[2012][332012] = People_332012
}
