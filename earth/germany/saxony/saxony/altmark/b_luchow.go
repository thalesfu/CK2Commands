package altmark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吕肖LuchowBarony struct {
	feud.BaseBarony
}

var BLuchow吕肖 feud.Barony = &吕肖LuchowBarony{}

func init() {
	f := BLuchow吕肖.(*吕肖LuchowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "luchow",
		TitleName: "吕肖",
		TitleCode: "b_luchow",
	}
}
