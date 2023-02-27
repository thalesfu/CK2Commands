package zeta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯卡达尔SkadarBarony struct {
	feud.BaseBarony
}

var BSkadar斯卡达尔 feud.Barony = &斯卡达尔SkadarBarony{}

func init() {
    f := BSkadar斯卡达尔.(*斯卡达尔SkadarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skadar",
		TitleName: "斯卡达尔",
		TitleCode: "b_skadar",
	}
}
