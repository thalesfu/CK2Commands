package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5200] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5200][125200] = People_125200
	HistoryPeopleMap[5200][145200] = People_145200
	HistoryPeopleMap[5200][155200] = People_155200
	HistoryPeopleMap[5200][205200] = People_205200
	HistoryPeopleMap[5200][235200] = People_235200
	HistoryPeopleMap[5200][475200] = People_475200
	HistoryPeopleMap[5200][485200] = People_485200
}
