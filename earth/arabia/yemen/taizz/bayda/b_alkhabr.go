package bayda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海卜尔AlkhabrBarony struct {
	feud.BaseBarony
}

var BAlkhabr海卜尔 feud.Barony = &海卜尔AlkhabrBarony{}

func init() {
    f := BAlkhabr海卜尔.(*海卜尔AlkhabrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alkhabr",
		TitleName: "海卜尔",
		TitleCode: "b_alkhabr",
	}
}
