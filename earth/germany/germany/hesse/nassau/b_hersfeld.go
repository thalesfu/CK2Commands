package nassau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫斯费尔德HersfeldBarony struct {
	feud.BaseBarony
}

var BHersfeld赫斯费尔德 feud.Barony = &赫斯费尔德HersfeldBarony{}

func init() {
	f := BHersfeld赫斯费尔德.(*赫斯费尔德HersfeldBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hersfeld",
		TitleName: "赫斯费尔德",
		TitleCode: "b_hersfeld",
	}
}
