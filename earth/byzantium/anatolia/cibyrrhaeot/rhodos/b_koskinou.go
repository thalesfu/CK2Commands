package rhodos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科斯基努KoskinouBarony struct {
	feud.BaseBarony
}

var BKoskinou科斯基努 feud.Barony = &科斯基努KoskinouBarony{}

func init() {
	f := BKoskinou科斯基努.(*科斯基努KoskinouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koskinou",
		TitleName: "科斯基努",
		TitleCode: "b_koskinou",
	}
}
