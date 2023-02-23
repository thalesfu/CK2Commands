package ani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔斯KarutsBarony struct {
	feud.BaseBarony
}

var BKaruts卡尔斯 feud.Barony = &卡尔斯KarutsBarony{}

func init() {
	f := BKaruts卡尔斯.(*卡尔斯KarutsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karuts",
		TitleName: "卡尔斯",
		TitleCode: "b_karuts",
	}
}
