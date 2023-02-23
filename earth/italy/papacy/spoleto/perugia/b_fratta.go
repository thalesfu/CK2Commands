package perugia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗拉塔FrattaBarony struct {
	feud.BaseBarony
}

var BFratta弗拉塔 feud.Barony = &弗拉塔FrattaBarony{}

func init() {
	f := BFratta弗拉塔.(*弗拉塔FrattaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fratta",
		TitleName: "弗拉塔",
		TitleCode: "b_fratta",
	}
}
