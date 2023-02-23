package vyazma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫赫纳特卡MokhnatkaBarony struct {
	feud.BaseBarony
}

var BMokhnatka莫赫纳特卡 feud.Barony = &莫赫纳特卡MokhnatkaBarony{}

func init() {
	f := BMokhnatka莫赫纳特卡.(*莫赫纳特卡MokhnatkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mokhnatka",
		TitleName: "莫赫纳特卡",
		TitleCode: "b_mokhnatka",
	}
}
