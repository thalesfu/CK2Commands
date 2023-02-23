package vasyugan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅吉翁MegionBarony struct {
	feud.BaseBarony
}

var BMegion梅吉翁 feud.Barony = &梅吉翁MegionBarony{}

func init() {
	f := BMegion梅吉翁.(*梅吉翁MegionBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "megion",
		TitleName: "梅吉翁",
		TitleCode: "b_megion",
	}
}
