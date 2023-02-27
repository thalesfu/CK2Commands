package ryazan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉斯诺耶KrasnoyeBarony struct {
	feud.BaseBarony
}

var BKrasnoye克拉斯诺耶 feud.Barony = &克拉斯诺耶KrasnoyeBarony{}

func init() {
    f := BKrasnoye克拉斯诺耶.(*克拉斯诺耶KrasnoyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krasnoye",
		TitleName: "克拉斯诺耶",
		TitleCode: "b_krasnoye",
	}
}
