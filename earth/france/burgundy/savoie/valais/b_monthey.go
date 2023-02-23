package valais

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙泰MontheyBarony struct {
	feud.BaseBarony
}

var BMonthey蒙泰 feud.Barony = &蒙泰MontheyBarony{}

func init() {
	f := BMonthey蒙泰.(*蒙泰MontheyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "monthey",
		TitleName: "蒙泰",
		TitleCode: "b_monthey",
	}
}
