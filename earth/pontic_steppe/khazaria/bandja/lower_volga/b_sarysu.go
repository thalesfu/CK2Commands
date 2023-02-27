package lower_volga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨雷苏SarysuBarony struct {
	feud.BaseBarony
}

var BSarysu萨雷苏 feud.Barony = &萨雷苏SarysuBarony{}

func init() {
    f := BSarysu萨雷苏.(*萨雷苏SarysuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarysu",
		TitleName: "萨雷苏",
		TitleCode: "b_sarysu",
	}
}
