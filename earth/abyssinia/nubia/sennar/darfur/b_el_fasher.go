package darfur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法希尔El_fasherBarony struct {
	feud.BaseBarony
}

var BEl_fasher法希尔 feud.Barony = &法希尔El_fasherBarony{}

func init() {
    f := BEl_fasher法希尔.(*法希尔El_fasherBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "el_fasher",
		TitleName: "法希尔",
		TitleCode: "b_el_fasher",
	}
}
