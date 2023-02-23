package braganza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布拉干萨BraganzaBarony struct {
	feud.BaseBarony
}

var BBraganza布拉干萨 feud.Barony = &布拉干萨BraganzaBarony{}

func init() {
	f := BBraganza布拉干萨.(*布拉干萨BraganzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "braganza",
		TitleName: "布拉干萨",
		TitleCode: "b_braganza",
	}
}
