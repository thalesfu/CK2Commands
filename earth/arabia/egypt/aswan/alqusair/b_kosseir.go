package alqusair

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库赛尔KosseirBarony struct {
	feud.BaseBarony
}

var BKosseir库赛尔 feud.Barony = &库赛尔KosseirBarony{}

func init() {
	f := BKosseir库赛尔.(*库赛尔KosseirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kosseir",
		TitleName: "库赛尔",
		TitleCode: "b_kosseir",
	}
}
