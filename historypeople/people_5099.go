package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5099] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5099][125099] = People_125099
	HistoryPeopleMap[5099][145099] = People_145099
	HistoryPeopleMap[5099][155099] = People_155099
	HistoryPeopleMap[5099][205099] = People_205099
	HistoryPeopleMap[5099][235099] = People_235099
}
