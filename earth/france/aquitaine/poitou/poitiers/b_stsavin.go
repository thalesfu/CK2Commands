package poitiers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣萨万StsavinBarony struct {
	feud.BaseBarony
}

var BStsavin圣萨万 feud.Barony = &圣萨万StsavinBarony{}

func init() {
	f := BStsavin圣萨万.(*圣萨万StsavinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stsavin",
		TitleName: "圣萨万",
		TitleCode: "b_stsavin",
	}
}
