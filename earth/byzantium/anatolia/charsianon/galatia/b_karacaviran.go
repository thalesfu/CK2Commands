package galatia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉贾维兰KaracaviranBarony struct {
	feud.BaseBarony
}

var BKaracaviran卡拉贾维兰 feud.Barony = &卡拉贾维兰KaracaviranBarony{}

func init() {
    f := BKaracaviran卡拉贾维兰.(*卡拉贾维兰KaracaviranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karacaviran",
		TitleName: "卡拉贾维兰",
		TitleCode: "b_karacaviran",
	}
}
