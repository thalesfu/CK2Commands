package holstein

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱因霍尔德堡ReinholdsburgBarony struct {
	feud.BaseBarony
}

var BReinholdsburg莱因霍尔德堡 feud.Barony = &莱因霍尔德堡ReinholdsburgBarony{}

func init() {
    f := BReinholdsburg莱因霍尔德堡.(*莱因霍尔德堡ReinholdsburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "reinholdsburg",
		TitleName: "莱因霍尔德堡",
		TitleCode: "b_reinholdsburg",
	}
}
