package taranto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫托拉MottolaBarony struct {
	feud.BaseBarony
}

var BMottola莫托拉 feud.Barony = &莫托拉MottolaBarony{}

func init() {
    f := BMottola莫托拉.(*莫托拉MottolaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mottola",
		TitleName: "莫托拉",
		TitleCode: "b_mottola",
	}
}
