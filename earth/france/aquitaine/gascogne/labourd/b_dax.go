package labourd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达克斯DaxBarony struct {
	feud.BaseBarony
}

var BDax达克斯 feud.Barony = &达克斯DaxBarony{}

func init() {
    f := BDax达克斯.(*达克斯DaxBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dax",
		TitleName: "达克斯",
		TitleCode: "b_dax",
	}
}
