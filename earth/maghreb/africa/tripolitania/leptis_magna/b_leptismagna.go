package leptis_magna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 大莱波蒂斯LeptismagnaBarony struct {
	feud.BaseBarony
}

var BLeptismagna大莱波蒂斯 feud.Barony = &大莱波蒂斯LeptismagnaBarony{}

func init() {
    f := BLeptismagna大莱波蒂斯.(*大莱波蒂斯LeptismagnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leptismagna",
		TitleName: "大莱波蒂斯",
		TitleCode: "b_leptismagna",
	}
}
