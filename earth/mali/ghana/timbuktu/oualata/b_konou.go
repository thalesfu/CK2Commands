package oualata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科努KonouBarony struct {
	feud.BaseBarony
}

var BKonou科努 feud.Barony = &科努KonouBarony{}

func init() {
    f := BKonou科努.(*科努KonouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "konou",
		TitleName: "科努",
		TitleCode: "b_konou",
	}
}
