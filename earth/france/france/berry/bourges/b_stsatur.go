package bourges

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣萨图尔StsaturBarony struct {
	feud.BaseBarony
}

var BStsatur圣萨图尔 feud.Barony = &圣萨图尔StsaturBarony{}

func init() {
	f := BStsatur圣萨图尔.(*圣萨图尔StsaturBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stsatur",
		TitleName: "圣萨图尔",
		TitleCode: "b_stsatur",
	}
}
