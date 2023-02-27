package vladimir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维亚兹尼基ViaznikiBarony struct {
	feud.BaseBarony
}

var BViazniki维亚兹尼基 feud.Barony = &维亚兹尼基ViaznikiBarony{}

func init() {
    f := BViazniki维亚兹尼基.(*维亚兹尼基ViaznikiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "viazniki",
		TitleName: "维亚兹尼基",
		TitleCode: "b_viazniki",
	}
}
