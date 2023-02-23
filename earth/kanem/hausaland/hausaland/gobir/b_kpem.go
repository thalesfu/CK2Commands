package gobir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩梅KpemBarony struct {
	feud.BaseBarony
}

var BKpem佩梅 feud.Barony = &佩梅KpemBarony{}

func init() {
	f := BKpem佩梅.(*佩梅KpemBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kpem",
		TitleName: "佩梅",
		TitleCode: "b_kpem",
	}
}
