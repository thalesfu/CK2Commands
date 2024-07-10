package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6800] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6800][6800] = People_6800
	HistoryPeopleMap[6800][166800] = People_166800
	HistoryPeopleMap[6800][206800] = People_206800
}
