package madaba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马德巴MadabaBarony struct {
	feud.BaseBarony
}

var BMadaba马德巴 feud.Barony = &马德巴MadabaBarony{}

func init() {
    f := BMadaba马德巴.(*马德巴MadabaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "madaba",
		TitleName: "马德巴",
		TitleCode: "b_madaba",
	}
}
