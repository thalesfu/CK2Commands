package lykia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利米拉LimyraBarony struct {
	feud.BaseBarony
}

var BLimyra利米拉 feud.Barony = &利米拉LimyraBarony{}

func init() {
    f := BLimyra利米拉.(*利米拉LimyraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "limyra",
		TitleName: "利米拉",
		TitleCode: "b_limyra",
	}
}
