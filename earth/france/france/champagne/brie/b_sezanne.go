package brie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞扎讷SezanneBarony struct {
	feud.BaseBarony
}

var BSezanne塞扎讷 feud.Barony = &塞扎讷SezanneBarony{}

func init() {
	f := BSezanne塞扎讷.(*塞扎讷SezanneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sezanne",
		TitleName: "塞扎讷",
		TitleCode: "b_sezanne",
	}
}
