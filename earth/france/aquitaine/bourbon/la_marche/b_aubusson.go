package la_marche

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欧比松AubussonBarony struct {
	feud.BaseBarony
}

var BAubusson欧比松 feud.Barony = &欧比松AubussonBarony{}

func init() {
    f := BAubusson欧比松.(*欧比松AubussonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aubusson",
		TitleName: "欧比松",
		TitleCode: "b_aubusson",
	}
}
