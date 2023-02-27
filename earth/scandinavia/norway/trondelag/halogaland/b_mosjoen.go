package halogaland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫舍恩MosjoenBarony struct {
	feud.BaseBarony
}

var BMosjoen莫舍恩 feud.Barony = &莫舍恩MosjoenBarony{}

func init() {
    f := BMosjoen莫舍恩.(*莫舍恩MosjoenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mosjoen",
		TitleName: "莫舍恩",
		TitleCode: "b_mosjoen",
	}
}
