package telemark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希特达拉斯HitterdalsBarony struct {
	feud.BaseBarony
}

var BHitterdals希特达拉斯 feud.Barony = &希特达拉斯HitterdalsBarony{}

func init() {
    f := BHitterdals希特达拉斯.(*希特达拉斯HitterdalsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hitterdals",
		TitleName: "希特达拉斯",
		TitleCode: "b_hitterdals",
	}
}
