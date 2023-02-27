package osnabruck

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维尔德斯豪森WildeshausenBarony struct {
	feud.BaseBarony
}

var BWildeshausen维尔德斯豪森 feud.Barony = &维尔德斯豪森WildeshausenBarony{}

func init() {
    f := BWildeshausen维尔德斯豪森.(*维尔德斯豪森WildeshausenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wildeshausen",
		TitleName: "维尔德斯豪森",
		TitleCode: "b_wildeshausen",
	}
}
