package cherven

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏捷伊斯克SuteyskBarony struct {
	feud.BaseBarony
}

var BSuteysk苏捷伊斯克 feud.Barony = &苏捷伊斯克SuteyskBarony{}

func init() {
    f := BSuteysk苏捷伊斯克.(*苏捷伊斯克SuteyskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suteysk",
		TitleName: "苏捷伊斯克",
		TitleCode: "b_suteysk",
	}
}
