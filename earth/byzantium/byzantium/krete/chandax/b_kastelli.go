package chandax

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯泰利KastelliBarony struct {
	feud.BaseBarony
}

var BKastelli卡斯泰利 feud.Barony = &卡斯泰利KastelliBarony{}

func init() {
    f := BKastelli卡斯泰利.(*卡斯泰利KastelliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kastelli",
		TitleName: "卡斯泰利",
		TitleCode: "b_kastelli",
	}
}
