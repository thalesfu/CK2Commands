package danzig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥利瓦OlivaBarony struct {
	feud.BaseBarony
}

var BOliva奥利瓦 feud.Barony = &奥利瓦OlivaBarony{}

func init() {
    f := BOliva奥利瓦.(*奥利瓦OlivaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oliva",
		TitleName: "奥利瓦",
		TitleCode: "b_oliva",
	}
}
