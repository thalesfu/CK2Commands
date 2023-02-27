package aj_bogd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 常村ChangcunBarony struct {
	feud.BaseBarony
}

var BChangcun常村 feud.Barony = &常村ChangcunBarony{}

func init() {
    f := BChangcun常村.(*常村ChangcunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "changcun",
		TitleName: "常村",
		TitleCode: "b_changcun",
	}
}
