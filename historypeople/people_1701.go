package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1701] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1701][1701] = People_1701
	HistoryPeopleMap[1701][31701] = People_31701
	HistoryPeopleMap[1701][41701] = People_41701
	HistoryPeopleMap[1701][71701] = People_71701
	HistoryPeopleMap[1701][131701] = People_131701
	HistoryPeopleMap[1701][161701] = People_161701
	HistoryPeopleMap[1701][191701] = People_191701
	HistoryPeopleMap[1701][221701] = People_221701
	HistoryPeopleMap[1701][461701] = People_461701
}
