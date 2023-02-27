package cieszyn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯科丘夫SkoczowBarony struct {
	feud.BaseBarony
}

var BSkoczow斯科丘夫 feud.Barony = &斯科丘夫SkoczowBarony{}

func init() {
    f := BSkoczow斯科丘夫.(*斯科丘夫SkoczowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skoczow",
		TitleName: "斯科丘夫",
		TitleCode: "b_skoczow",
	}
}
