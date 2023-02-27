package al_amarah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿马拉AmarahBarony struct {
	feud.BaseBarony
}

var BAmarah阿马拉 feud.Barony = &阿马拉AmarahBarony{}

func init() {
    f := BAmarah阿马拉.(*阿马拉AmarahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amarah",
		TitleName: "阿马拉",
		TitleCode: "b_amarah",
	}
}
