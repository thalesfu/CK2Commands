package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6001] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6001][6001] = People_6001
	HistoryPeopleMap[6001][96001] = People_96001
	HistoryPeopleMap[6001][126001] = People_126001
	HistoryPeopleMap[6001][146001] = People_146001
	HistoryPeopleMap[6001][166001] = People_166001
	HistoryPeopleMap[6001][186001] = People_186001
	HistoryPeopleMap[6001][206001] = People_206001
	HistoryPeopleMap[6001][216001] = People_216001
	HistoryPeopleMap[6001][256001] = People_256001
}
