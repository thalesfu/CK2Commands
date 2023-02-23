package cadiz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣何塞斯德尔巴列SanjosedelvalleBarony struct {
	feud.BaseBarony
}

var BSanjosedelvalle圣何塞斯德尔巴列 feud.Barony = &圣何塞斯德尔巴列SanjosedelvalleBarony{}

func init() {
	f := BSanjosedelvalle圣何塞斯德尔巴列.(*圣何塞斯德尔巴列SanjosedelvalleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanjosedelvalle",
		TitleName: "圣何塞斯德尔巴列",
		TitleCode: "b_sanjosedelvalle",
	}
}
