package bamberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡多尔茨堡CadolzburgBarony struct {
	feud.BaseBarony
}

var BCadolzburg卡多尔茨堡 feud.Barony = &卡多尔茨堡CadolzburgBarony{}

func init() {
	f := BCadolzburg卡多尔茨堡.(*卡多尔茨堡CadolzburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cadolzburg",
		TitleName: "卡多尔茨堡",
		TitleCode: "b_cadolzburg",
	}
}
