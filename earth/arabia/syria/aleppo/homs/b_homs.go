package homs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍姆斯HomsBarony struct {
	feud.BaseBarony
}

var BHoms霍姆斯 feud.Barony = &霍姆斯HomsBarony{}

func init() {
    f := BHoms霍姆斯.(*霍姆斯HomsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "homs",
		TitleName: "霍姆斯",
		TitleCode: "b_homs",
	}
}
