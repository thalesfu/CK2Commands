package orvieto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴斯基BaschiBarony struct {
	feud.BaseBarony
}

var BBaschi巴斯基 feud.Barony = &巴斯基BaschiBarony{}

func init() {
    f := BBaschi巴斯基.(*巴斯基BaschiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baschi",
		TitleName: "巴斯基",
		TitleCode: "b_baschi",
	}
}
