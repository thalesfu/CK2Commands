package bolzano

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅兰MeranBarony struct {
	feud.BaseBarony
}

var BMeran梅兰 feud.Barony = &梅兰MeranBarony{}

func init() {
	f := BMeran梅兰.(*梅兰MeranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "meran",
		TitleName: "梅兰",
		TitleCode: "b_meran",
	}
}
