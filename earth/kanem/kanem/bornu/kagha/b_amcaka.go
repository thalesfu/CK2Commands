package kagha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿姆察卡AmcakaBarony struct {
	feud.BaseBarony
}

var BAmcaka阿姆察卡 feud.Barony = &阿姆察卡AmcakaBarony{}

func init() {
	f := BAmcaka阿姆察卡.(*阿姆察卡AmcakaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amcaka",
		TitleName: "阿姆察卡",
		TitleCode: "b_amcaka",
	}
}
