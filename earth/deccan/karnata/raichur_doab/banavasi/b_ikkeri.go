package banavasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊凯里IkkeriBarony struct {
	feud.BaseBarony
}

var BIkkeri伊凯里 feud.Barony = &伊凯里IkkeriBarony{}

func init() {
    f := BIkkeri伊凯里.(*伊凯里IkkeriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ikkeri",
		TitleName: "伊凯里",
		TitleCode: "b_ikkeri",
	}
}
