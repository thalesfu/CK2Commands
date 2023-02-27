package pressburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫多尔ModorBarony struct {
	feud.BaseBarony
}

var BModor莫多尔 feud.Barony = &莫多尔ModorBarony{}

func init() {
    f := BModor莫多尔.(*莫多尔ModorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "modor",
		TitleName: "莫多尔",
		TitleCode: "b_modor",
	}
}
