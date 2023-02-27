package tsagaannuur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科布多KhovdBarony struct {
	feud.BaseBarony
}

var BKhovd科布多 feud.Barony = &科布多KhovdBarony{}

func init() {
    f := BKhovd科布多.(*科布多KhovdBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khovd",
		TitleName: "科布多",
		TitleCode: "b_khovd",
	}
}
