package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5098] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5098][125098] = People_125098
	HistoryPeopleMap[5098][145098] = People_145098
	HistoryPeopleMap[5098][155098] = People_155098
	HistoryPeopleMap[5098][205098] = People_205098
	HistoryPeopleMap[5098][235098] = People_235098
}
