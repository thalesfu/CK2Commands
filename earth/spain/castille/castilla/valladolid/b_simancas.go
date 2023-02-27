package valladolid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡曼卡斯SimancasBarony struct {
	feud.BaseBarony
}

var BSimancas锡曼卡斯 feud.Barony = &锡曼卡斯SimancasBarony{}

func init() {
    f := BSimancas锡曼卡斯.(*锡曼卡斯SimancasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "simancas",
		TitleName: "锡曼卡斯",
		TitleCode: "b_simancas",
	}
}
