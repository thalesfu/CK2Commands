package bukhara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔库赫KarkuhBarony struct {
	feud.BaseBarony
}

var BKarkuh卡尔库赫 feud.Barony = &卡尔库赫KarkuhBarony{}

func init() {
	f := BKarkuh卡尔库赫.(*卡尔库赫KarkuhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karkuh",
		TitleName: "卡尔库赫",
		TitleCode: "b_karkuh",
	}
}
