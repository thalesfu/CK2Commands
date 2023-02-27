package pavia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯泰焦CasteggioBarony struct {
	feud.BaseBarony
}

var BCasteggio卡斯泰焦 feud.Barony = &卡斯泰焦CasteggioBarony{}

func init() {
    f := BCasteggio卡斯泰焦.(*卡斯泰焦CasteggioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "casteggio",
		TitleName: "卡斯泰焦",
		TitleCode: "b_casteggio",
	}
}
