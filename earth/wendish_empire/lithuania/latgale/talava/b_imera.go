package talava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊梅拉ImeraBarony struct {
	feud.BaseBarony
}

var BImera伊梅拉 feud.Barony = &伊梅拉ImeraBarony{}

func init() {
    f := BImera伊梅拉.(*伊梅拉ImeraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "imera",
		TitleName: "伊梅拉",
		TitleCode: "b_imera",
	}
}
