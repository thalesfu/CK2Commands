package sibir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯克KaikBarony struct {
	feud.BaseBarony
}

var BKaik凯克 feud.Barony = &凯克KaikBarony{}

func init() {
    f := BKaik凯克.(*凯克KaikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaik",
		TitleName: "凯克",
		TitleCode: "b_kaik",
	}
}
