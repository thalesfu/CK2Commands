package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6789] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6789][166789] = People_166789
	HistoryPeopleMap[6789][206789] = People_206789
}
