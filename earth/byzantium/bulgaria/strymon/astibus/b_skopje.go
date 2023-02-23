package astibus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯科普里SkopjeBarony struct {
	feud.BaseBarony
}

var BSkopje斯科普里 feud.Barony = &斯科普里SkopjeBarony{}

func init() {
	f := BSkopje斯科普里.(*斯科普里SkopjeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skopje",
		TitleName: "斯科普里",
		TitleCode: "b_skopje",
	}
}
