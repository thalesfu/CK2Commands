package severin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔绍瓦OrsovaBarony struct {
	feud.BaseBarony
}

var BOrsova奥尔绍瓦 feud.Barony = &奥尔绍瓦OrsovaBarony{}

func init() {
    f := BOrsova奥尔绍瓦.(*奥尔绍瓦OrsovaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orsova",
		TitleName: "奥尔绍瓦",
		TitleCode: "b_orsova",
	}
}
