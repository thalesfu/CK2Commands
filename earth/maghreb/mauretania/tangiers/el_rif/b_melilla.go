package el_rif

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅利利亚MelillaBarony struct {
	feud.BaseBarony
}

var BMelilla梅利利亚 feud.Barony = &梅利利亚MelillaBarony{}

func init() {
    f := BMelilla梅利利亚.(*梅利利亚MelillaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "melilla",
		TitleName: "梅利利亚",
		TitleCode: "b_melilla",
	}
}
