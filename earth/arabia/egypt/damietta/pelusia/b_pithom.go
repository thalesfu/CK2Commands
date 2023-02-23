package pelusia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮索姆PithomBarony struct {
	feud.BaseBarony
}

var BPithom皮索姆 feud.Barony = &皮索姆PithomBarony{}

func init() {
	f := BPithom皮索姆.(*皮索姆PithomBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pithom",
		TitleName: "皮索姆",
		TitleCode: "b_pithom",
	}
}
