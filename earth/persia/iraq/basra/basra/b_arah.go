package basra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉ArahBarony struct {
	feud.BaseBarony
}

var BArah阿拉 feud.Barony = &阿拉ArahBarony{}

func init() {
    f := BArah阿拉.(*阿拉ArahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arah",
		TitleName: "阿拉",
		TitleCode: "b_arah",
	}
}
