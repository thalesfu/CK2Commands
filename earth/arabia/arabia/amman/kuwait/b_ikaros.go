package kuwait

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊卡洛斯IkarosBarony struct {
	feud.BaseBarony
}

var BIkaros伊卡洛斯 feud.Barony = &伊卡洛斯IkarosBarony{}

func init() {
    f := BIkaros伊卡洛斯.(*伊卡洛斯IkarosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ikaros",
		TitleName: "伊卡洛斯",
		TitleCode: "b_ikaros",
	}
}
