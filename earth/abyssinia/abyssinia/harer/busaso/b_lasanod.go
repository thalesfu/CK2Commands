package busaso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉斯阿诺得LasanodBarony struct {
	feud.BaseBarony
}

var BLasanod拉斯阿诺得 feud.Barony = &拉斯阿诺得LasanodBarony{}

func init() {
    f := BLasanod拉斯阿诺得.(*拉斯阿诺得LasanodBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lasanod",
		TitleName: "拉斯阿诺得",
		TitleCode: "b_lasanod",
	}
}
