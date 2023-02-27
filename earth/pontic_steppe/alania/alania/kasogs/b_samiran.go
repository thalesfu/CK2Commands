package kasogs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨米兰SamiranBarony struct {
	feud.BaseBarony
}

var BSamiran萨米兰 feud.Barony = &萨米兰SamiranBarony{}

func init() {
    f := BSamiran萨米兰.(*萨米兰SamiranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samiran",
		TitleName: "萨米兰",
		TitleCode: "b_samiran",
	}
}
