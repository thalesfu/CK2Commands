package aydhab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈拉伊卜HalayebBarony struct {
	feud.BaseBarony
}

var BHalayeb哈拉伊卜 feud.Barony = &哈拉伊卜HalayebBarony{}

func init() {
    f := BHalayeb哈拉伊卜.(*哈拉伊卜HalayebBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "halayeb",
		TitleName: "哈拉伊卜",
		TitleCode: "b_halayeb",
	}
}
