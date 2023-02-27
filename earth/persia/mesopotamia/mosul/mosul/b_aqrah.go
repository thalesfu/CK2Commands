package mosul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿奇拉AqrahBarony struct {
	feud.BaseBarony
}

var BAqrah阿奇拉 feud.Barony = &阿奇拉AqrahBarony{}

func init() {
    f := BAqrah阿奇拉.(*阿奇拉AqrahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aqrah",
		TitleName: "阿奇拉",
		TitleCode: "b_aqrah",
	}
}
