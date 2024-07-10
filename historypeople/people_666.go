package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[666] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[666][666] = People_666
	HistoryPeopleMap[666][20666] = People_20666
	HistoryPeopleMap[666][70666] = People_70666
	HistoryPeopleMap[666][180666] = People_180666
	HistoryPeopleMap[666][260666] = People_260666
}
