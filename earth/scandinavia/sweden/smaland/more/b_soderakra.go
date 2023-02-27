package more

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索德罗克拉SoderakraBarony struct {
	feud.BaseBarony
}

var BSoderakra索德罗克拉 feud.Barony = &索德罗克拉SoderakraBarony{}

func init() {
    f := BSoderakra索德罗克拉.(*索德罗克拉SoderakraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "soderakra",
		TitleName: "索德罗克拉",
		TitleCode: "b_soderakra",
	}
}
