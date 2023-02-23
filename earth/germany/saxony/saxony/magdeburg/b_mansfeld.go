package magdeburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼斯费尔德MansfeldBarony struct {
	feud.BaseBarony
}

var BMansfeld曼斯费尔德 feud.Barony = &曼斯费尔德MansfeldBarony{}

func init() {
	f := BMansfeld曼斯费尔德.(*曼斯费尔德MansfeldBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mansfeld",
		TitleName: "曼斯费尔德",
		TitleCode: "b_mansfeld",
	}
}
