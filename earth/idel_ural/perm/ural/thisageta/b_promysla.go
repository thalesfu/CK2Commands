package thisageta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普罗梅斯拉PromyslaBarony struct {
	feud.BaseBarony
}

var BPromysla普罗梅斯拉 feud.Barony = &普罗梅斯拉PromyslaBarony{}

func init() {
    f := BPromysla普罗梅斯拉.(*普罗梅斯拉PromyslaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "promysla",
		TitleName: "普罗梅斯拉",
		TitleCode: "b_promysla",
	}
}
