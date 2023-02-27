package kasogs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨雷_秋兹SarytyuzBarony struct {
	feud.BaseBarony
}

var BSarytyuz萨雷_秋兹 feud.Barony = &萨雷_秋兹SarytyuzBarony{}

func init() {
    f := BSarytyuz萨雷_秋兹.(*萨雷_秋兹SarytyuzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarytyuz",
		TitleName: "萨雷_秋兹",
		TitleCode: "b_sarytyuz",
	}
}
