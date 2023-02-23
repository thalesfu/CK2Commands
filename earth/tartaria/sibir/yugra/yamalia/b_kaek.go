package yamalia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡耶克KaekBarony struct {
	feud.BaseBarony
}

var BKaek卡耶克 feud.Barony = &卡耶克KaekBarony{}

func init() {
	f := BKaek卡耶克.(*卡耶克KaekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaek",
		TitleName: "卡耶克",
		TitleCode: "b_kaek",
	}
}
