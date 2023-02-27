package thessalia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉里萨LarissaBarony struct {
	feud.BaseBarony
}

var BLarissa拉里萨 feud.Barony = &拉里萨LarissaBarony{}

func init() {
    f := BLarissa拉里萨.(*拉里萨LarissaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "larissa",
		TitleName: "拉里萨",
		TitleCode: "b_larissa",
	}
}
