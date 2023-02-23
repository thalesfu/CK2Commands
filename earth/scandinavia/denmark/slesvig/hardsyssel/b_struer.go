package hardsyssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯楚厄StruerBarony struct {
	feud.BaseBarony
}

var BStruer斯楚厄 feud.Barony = &斯楚厄StruerBarony{}

func init() {
	f := BStruer斯楚厄.(*斯楚厄StruerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "struer",
		TitleName: "斯楚厄",
		TitleCode: "b_struer",
	}
}
