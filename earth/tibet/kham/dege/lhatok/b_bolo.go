package lhatok

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫洛BoloBarony struct {
	feud.BaseBarony
}

var BBolo莫洛 feud.Barony = &莫洛BoloBarony{}

func init() {
    f := BBolo莫洛.(*莫洛BoloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bolo",
		TitleName: "莫洛",
		TitleCode: "b_bolo",
	}
}
