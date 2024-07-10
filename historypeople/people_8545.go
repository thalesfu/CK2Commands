package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8545] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8545][168545] = People_168545
	HistoryPeopleMap[8545][188545] = People_188545
}
