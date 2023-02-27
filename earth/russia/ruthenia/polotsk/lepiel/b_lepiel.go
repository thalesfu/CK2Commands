package lepiel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 列佩利LepielBarony struct {
	feud.BaseBarony
}

var BLepiel列佩利 feud.Barony = &列佩利LepielBarony{}

func init() {
    f := BLepiel列佩利.(*列佩利LepielBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lepiel",
		TitleName: "列佩利",
		TitleCode: "b_lepiel",
	}
}
