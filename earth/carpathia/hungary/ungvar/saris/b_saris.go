package saris

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯皮什SarisBarony struct {
	feud.BaseBarony
}

var BSaris斯皮什 feud.Barony = &斯皮什SarisBarony{}

func init() {
	f := BSaris斯皮什.(*斯皮什SarisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saris",
		TitleName: "斯皮什",
		TitleCode: "b_saris",
	}
}
