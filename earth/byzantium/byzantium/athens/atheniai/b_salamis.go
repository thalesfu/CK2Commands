package atheniai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨拉米斯SalamisBarony struct {
	feud.BaseBarony
}

var BSalamis萨拉米斯 feud.Barony = &萨拉米斯SalamisBarony{}

func init() {
    f := BSalamis萨拉米斯.(*萨拉米斯SalamisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salamis",
		TitleName: "萨拉米斯",
		TitleCode: "b_salamis",
	}
}
