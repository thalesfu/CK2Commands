package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1905] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1905][71905] = People_71905
	HistoryPeopleMap[1905][161905] = People_161905
	HistoryPeopleMap[1905][191905] = People_191905
	HistoryPeopleMap[1905][461905] = People_461905
}
