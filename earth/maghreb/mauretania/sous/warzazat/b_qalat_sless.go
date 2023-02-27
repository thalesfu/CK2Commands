package warzazat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯莱斯堡Qalat_slessBarony struct {
	feud.BaseBarony
}

var BQalat_sless斯莱斯堡 feud.Barony = &斯莱斯堡Qalat_slessBarony{}

func init() {
    f := BQalat_sless斯莱斯堡.(*斯莱斯堡Qalat_slessBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qalat_sless",
		TitleName: "斯莱斯堡",
		TitleCode: "b_qalat_sless",
	}
}
