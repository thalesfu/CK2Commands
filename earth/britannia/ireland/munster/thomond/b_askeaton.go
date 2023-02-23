package thomond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿斯基顿AskeatonBarony struct {
	feud.BaseBarony
}

var BAskeaton阿斯基顿 feud.Barony = &阿斯基顿AskeatonBarony{}

func init() {
	f := BAskeaton阿斯基顿.(*阿斯基顿AskeatonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "askeaton",
		TitleName: "阿斯基顿",
		TitleCode: "b_askeaton",
	}
}
