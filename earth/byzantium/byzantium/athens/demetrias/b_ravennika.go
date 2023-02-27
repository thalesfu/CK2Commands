package demetrias

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉文尼卡RavennikaBarony struct {
	feud.BaseBarony
}

var BRavennika拉文尼卡 feud.Barony = &拉文尼卡RavennikaBarony{}

func init() {
    f := BRavennika拉文尼卡.(*拉文尼卡RavennikaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ravennika",
		TitleName: "拉文尼卡",
		TitleCode: "b_ravennika",
	}
}
