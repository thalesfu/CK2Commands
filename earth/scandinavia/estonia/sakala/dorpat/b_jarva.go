package dorpat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶尔瓦JarvaBarony struct {
	feud.BaseBarony
}

var BJarva耶尔瓦 feud.Barony = &耶尔瓦JarvaBarony{}

func init() {
    f := BJarva耶尔瓦.(*耶尔瓦JarvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jarva",
		TitleName: "耶尔瓦",
		TitleCode: "b_jarva",
	}
}
