package finland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 楠塔利NaantaliBarony struct {
	feud.BaseBarony
}

var BNaantali楠塔利 feud.Barony = &楠塔利NaantaliBarony{}

func init() {
    f := BNaantali楠塔利.(*楠塔利NaantaliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naantali",
		TitleName: "楠塔利",
		TitleCode: "b_naantali",
	}
}
