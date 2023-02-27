package socotra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加迪卜QadubBarony struct {
	feud.BaseBarony
}

var BQadub加迪卜 feud.Barony = &加迪卜QadubBarony{}

func init() {
    f := BQadub加迪卜.(*加迪卜QadubBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qadub",
		TitleName: "加迪卜",
		TitleCode: "b_qadub",
	}
}
