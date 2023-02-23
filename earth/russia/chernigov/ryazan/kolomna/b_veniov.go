package kolomna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦尼奥夫VeniovBarony struct {
	feud.BaseBarony
}

var BVeniov韦尼奥夫 feud.Barony = &韦尼奥夫VeniovBarony{}

func init() {
	f := BVeniov韦尼奥夫.(*韦尼奥夫VeniovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "veniov",
		TitleName: "韦尼奥夫",
		TitleCode: "b_veniov",
	}
}
