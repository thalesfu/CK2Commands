package rylsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科列诺沃KorenovoBarony struct {
	feud.BaseBarony
}

var BKorenovo科列诺沃 feud.Barony = &科列诺沃KorenovoBarony{}

func init() {
    f := BKorenovo科列诺沃.(*科列诺沃KorenovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "korenovo",
		TitleName: "科列诺沃",
		TitleCode: "b_korenovo",
	}
}
