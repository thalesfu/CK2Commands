package jarva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦伊加VaigaBarony struct {
	feud.BaseBarony
}

var BVaiga瓦伊加 feud.Barony = &瓦伊加VaigaBarony{}

func init() {
	f := BVaiga瓦伊加.(*瓦伊加VaigaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vaiga",
		TitleName: "瓦伊加",
		TitleCode: "b_vaiga",
	}
}
