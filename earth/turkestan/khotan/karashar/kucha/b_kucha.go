package kucha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苦叉KuchaBarony struct {
	feud.BaseBarony
}

var BKucha苦叉 feud.Barony = &苦叉KuchaBarony{}

func init() {
	f := BKucha苦叉.(*苦叉KuchaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kucha",
		TitleName: "苦叉",
		TitleCode: "b_kucha",
	}
}
