package mohadavasaka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 檀波DantaBarony struct {
	feud.BaseBarony
}

var BDanta檀波 feud.Barony = &檀波DantaBarony{}

func init() {
    f := BDanta檀波.(*檀波DantaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "danta",
		TitleName: "檀波",
		TitleCode: "b_danta",
	}
}
