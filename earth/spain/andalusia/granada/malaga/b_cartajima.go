package malaga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔塔希马CartajimaBarony struct {
	feud.BaseBarony
}

var BCartajima卡尔塔希马 feud.Barony = &卡尔塔希马CartajimaBarony{}

func init() {
    f := BCartajima卡尔塔希马.(*卡尔塔希马CartajimaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cartajima",
		TitleName: "卡尔塔希马",
		TitleCode: "b_cartajima",
	}
}
