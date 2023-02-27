package alabuga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫日加MozhgaBarony struct {
	feud.BaseBarony
}

var BMozhga莫日加 feud.Barony = &莫日加MozhgaBarony{}

func init() {
    f := BMozhga莫日加.(*莫日加MozhgaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mozhga",
		TitleName: "莫日加",
		TitleCode: "b_mozhga",
	}
}
