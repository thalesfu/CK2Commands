package uvek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯梅洛夫卡SmelovkaBarony struct {
	feud.BaseBarony
}

var BSmelovka斯梅洛夫卡 feud.Barony = &斯梅洛夫卡SmelovkaBarony{}

func init() {
    f := BSmelovka斯梅洛夫卡.(*斯梅洛夫卡SmelovkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "smelovka",
		TitleName: "斯梅洛夫卡",
		TitleCode: "b_smelovka",
	}
}
