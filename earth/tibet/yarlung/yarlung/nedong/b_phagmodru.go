package nedong

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伯木古鲁PhagmodruBarony struct {
	feud.BaseBarony
}

var BPhagmodru伯木古鲁 feud.Barony = &伯木古鲁PhagmodruBarony{}

func init() {
    f := BPhagmodru伯木古鲁.(*伯木古鲁PhagmodruBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "phagmodru",
		TitleName: "伯木古鲁",
		TitleCode: "b_phagmodru",
	}
}
