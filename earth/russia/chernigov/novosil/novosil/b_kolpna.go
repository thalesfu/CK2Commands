package novosil

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔普纳KolpnaBarony struct {
	feud.BaseBarony
}

var BKolpna科尔普纳 feud.Barony = &科尔普纳KolpnaBarony{}

func init() {
    f := BKolpna科尔普纳.(*科尔普纳KolpnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kolpna",
		TitleName: "科尔普纳",
		TitleCode: "b_kolpna",
	}
}
