package burhanpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶瓦YawalBarony struct {
	feud.BaseBarony
}

var BYawal耶瓦 feud.Barony = &耶瓦YawalBarony{}

func init() {
    f := BYawal耶瓦.(*耶瓦YawalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yawal",
		TitleName: "耶瓦",
		TitleCode: "b_yawal",
	}
}
