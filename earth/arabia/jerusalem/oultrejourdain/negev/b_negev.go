package negev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 内盖夫NegevBarony struct {
	feud.BaseBarony
}

var BNegev内盖夫 feud.Barony = &内盖夫NegevBarony{}

func init() {
    f := BNegev内盖夫.(*内盖夫NegevBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "negev",
		TitleName: "内盖夫",
		TitleCode: "b_negev",
	}
}
