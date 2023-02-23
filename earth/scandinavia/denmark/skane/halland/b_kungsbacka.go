package halland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孔斯巴卡KungsbackaBarony struct {
	feud.BaseBarony
}

var BKungsbacka孔斯巴卡 feud.Barony = &孔斯巴卡KungsbackaBarony{}

func init() {
	f := BKungsbacka孔斯巴卡.(*孔斯巴卡KungsbackaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kungsbacka",
		TitleName: "孔斯巴卡",
		TitleCode: "b_kungsbacka",
	}
}
