package travunia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德拉切维察DracevicaBarony struct {
	feud.BaseBarony
}

var BDracevica德拉切维察 feud.Barony = &德拉切维察DracevicaBarony{}

func init() {
	f := BDracevica德拉切维察.(*德拉切维察DracevicaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dracevica",
		TitleName: "德拉切维察",
		TitleCode: "b_dracevica",
	}
}
