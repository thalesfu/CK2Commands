package zavarot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥克西诺OksinoBarony struct {
	feud.BaseBarony
}

var BOksino奥克西诺 feud.Barony = &奥克西诺OksinoBarony{}

func init() {
    f := BOksino奥克西诺.(*奥克西诺OksinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oksino",
		TitleName: "奥克西诺",
		TitleCode: "b_oksino",
	}
}
