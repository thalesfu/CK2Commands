package bulgar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 厄克IqBarony struct {
	feud.BaseBarony
}

var BIq厄克 feud.Barony = &厄克IqBarony{}

func init() {
    f := BIq厄克.(*厄克IqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "iq",
		TitleName: "厄克",
		TitleCode: "b_iq",
	}
}
