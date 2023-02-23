package finland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库西斯托KuusistoBarony struct {
	feud.BaseBarony
}

var BKuusisto库西斯托 feud.Barony = &库西斯托KuusistoBarony{}

func init() {
	f := BKuusisto库西斯托.(*库西斯托KuusistoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuusisto",
		TitleName: "库西斯托",
		TitleCode: "b_kuusisto",
	}
}
