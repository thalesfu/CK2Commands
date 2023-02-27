package saravan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉哈雷PahrahBarony struct {
	feud.BaseBarony
}

var BPahrah拉哈雷 feud.Barony = &拉哈雷PahrahBarony{}

func init() {
    f := BPahrah拉哈雷.(*拉哈雷PahrahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pahrah",
		TitleName: "拉哈雷",
		TitleCode: "b_pahrah",
	}
}
