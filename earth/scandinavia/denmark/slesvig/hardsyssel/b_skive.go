package hardsyssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯基沃SkiveBarony struct {
	feud.BaseBarony
}

var BSkive斯基沃 feud.Barony = &斯基沃SkiveBarony{}

func init() {
	f := BSkive斯基沃.(*斯基沃SkiveBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skive",
		TitleName: "斯基沃",
		TitleCode: "b_skive",
	}
}
