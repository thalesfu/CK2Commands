package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1337] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1337][1337] = People_1337
	HistoryPeopleMap[1337][31337] = People_31337
	HistoryPeopleMap[1337][41337] = People_41337
	HistoryPeopleMap[1337][71337] = People_71337
	HistoryPeopleMap[1337][91337] = People_91337
	HistoryPeopleMap[1337][161337] = People_161337
	HistoryPeopleMap[1337][191337] = People_191337
	HistoryPeopleMap[1337][261337] = People_261337
}
