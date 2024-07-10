package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2015] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2015][2015] = People_2015
	HistoryPeopleMap[2015][12015] = People_12015
	HistoryPeopleMap[2015][42015] = People_42015
	HistoryPeopleMap[2015][72015] = People_72015
	HistoryPeopleMap[2015][82015] = People_82015
	HistoryPeopleMap[2015][142015] = People_142015
	HistoryPeopleMap[2015][182015] = People_182015
	HistoryPeopleMap[2015][192015] = People_192015
	HistoryPeopleMap[2015][212015] = People_212015
	HistoryPeopleMap[2015][232015] = People_232015
	HistoryPeopleMap[2015][252015] = People_252015
	HistoryPeopleMap[2015][332015] = People_332015
}
