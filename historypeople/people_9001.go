package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9001] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9001][129001] = People_129001
	HistoryPeopleMap[9001][159001] = People_159001
	HistoryPeopleMap[9001][189001] = People_189001
	HistoryPeopleMap[9001][219001] = People_219001
	HistoryPeopleMap[9001][229001] = People_229001
	HistoryPeopleMap[9001][259001] = People_259001
}
