package charsianon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德治文DazimonBarony struct {
	feud.BaseBarony
}

var BDazimon德治文 feud.Barony = &德治文DazimonBarony{}

func init() {
    f := BDazimon德治文.(*德治文DazimonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dazimon",
		TitleName: "德治文",
		TitleCode: "b_dazimon",
	}
}
