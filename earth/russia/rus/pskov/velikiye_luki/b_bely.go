package velikiye_luki

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别雷BelyBarony struct {
	feud.BaseBarony
}

var BBely别雷 feud.Barony = &别雷BelyBarony{}

func init() {
    f := BBely别雷.(*别雷BelyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bely",
		TitleName: "别雷",
		TitleCode: "b_bely",
	}
}
