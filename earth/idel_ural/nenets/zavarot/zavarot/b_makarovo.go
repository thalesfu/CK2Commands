package zavarot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马卡罗沃MakarovoBarony struct {
	feud.BaseBarony
}

var BMakarovo马卡罗沃 feud.Barony = &马卡罗沃MakarovoBarony{}

func init() {
    f := BMakarovo马卡罗沃.(*马卡罗沃MakarovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "makarovo",
		TitleName: "马卡罗沃",
		TitleCode: "b_makarovo",
	}
}
