package danzig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普茨克PuckBarony struct {
	feud.BaseBarony
}

var BPuck普茨克 feud.Barony = &普茨克PuckBarony{}

func init() {
    f := BPuck普茨克.(*普茨克PuckBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "puck",
		TitleName: "普茨克",
		TitleCode: "b_puck",
	}
}
