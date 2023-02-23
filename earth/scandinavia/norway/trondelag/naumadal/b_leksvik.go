package naumadal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱克斯维克LeksvikBarony struct {
	feud.BaseBarony
}

var BLeksvik莱克斯维克 feud.Barony = &莱克斯维克LeksvikBarony{}

func init() {
	f := BLeksvik莱克斯维克.(*莱克斯维克LeksvikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leksvik",
		TitleName: "莱克斯维克",
		TitleCode: "b_leksvik",
	}
}
