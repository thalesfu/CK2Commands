package thuringen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾希斯费尔德EichsfeldBarony struct {
	feud.BaseBarony
}

var BEichsfeld艾希斯费尔德 feud.Barony = &艾希斯费尔德EichsfeldBarony{}

func init() {
	f := BEichsfeld艾希斯费尔德.(*艾希斯费尔德EichsfeldBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eichsfeld",
		TitleName: "艾希斯费尔德",
		TitleCode: "b_eichsfeld",
	}
}
