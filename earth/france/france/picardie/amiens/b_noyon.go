package amiens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 努瓦永NoyonBarony struct {
	feud.BaseBarony
}

var BNoyon努瓦永 feud.Barony = &努瓦永NoyonBarony{}

func init() {
    f := BNoyon努瓦永.(*努瓦永NoyonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "noyon",
		TitleName: "努瓦永",
		TitleCode: "b_noyon",
	}
}
