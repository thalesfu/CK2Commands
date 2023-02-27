package bytow

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱贝LebeBarony struct {
	feud.BaseBarony
}

var BLebe莱贝 feud.Barony = &莱贝LebeBarony{}

func init() {
    f := BLebe莱贝.(*莱贝LebeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lebe",
		TitleName: "莱贝",
		TitleCode: "b_lebe",
	}
}
