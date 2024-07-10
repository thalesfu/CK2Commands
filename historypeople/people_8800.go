package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8800] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8800][168800] = People_168800
	HistoryPeopleMap[8800][188800] = People_188800
	HistoryPeopleMap[8800][218800] = People_218800
	HistoryPeopleMap[8800][248800] = People_248800
	HistoryPeopleMap[8800][468800] = People_468800
}
