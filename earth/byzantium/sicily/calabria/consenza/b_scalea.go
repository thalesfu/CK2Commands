package consenza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯卡莱阿ScaleaBarony struct {
	feud.BaseBarony
}

var BScalea斯卡莱阿 feud.Barony = &斯卡莱阿ScaleaBarony{}

func init() {
    f := BScalea斯卡莱阿.(*斯卡莱阿ScaleaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "scalea",
		TitleName: "斯卡莱阿",
		TitleCode: "b_scalea",
	}
}
