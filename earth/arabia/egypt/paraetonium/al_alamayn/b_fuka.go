package al_alamayn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 富凯FukaBarony struct {
	feud.BaseBarony
}

var BFuka富凯 feud.Barony = &富凯FukaBarony{}

func init() {
    f := BFuka富凯.(*富凯FukaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fuka",
		TitleName: "富凯",
		TitleCode: "b_fuka",
	}
}
