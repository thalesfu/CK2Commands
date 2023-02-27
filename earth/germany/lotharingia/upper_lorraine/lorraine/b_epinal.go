package lorraine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃皮纳勒EpinalBarony struct {
	feud.BaseBarony
}

var BEpinal埃皮纳勒 feud.Barony = &埃皮纳勒EpinalBarony{}

func init() {
    f := BEpinal埃皮纳勒.(*埃皮纳勒EpinalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "epinal",
		TitleName: "埃皮纳勒",
		TitleCode: "b_epinal",
	}
}
