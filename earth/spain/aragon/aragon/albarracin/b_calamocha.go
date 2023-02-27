package albarracin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉莫查CalamochaBarony struct {
	feud.BaseBarony
}

var BCalamocha卡拉莫查 feud.Barony = &卡拉莫查CalamochaBarony{}

func init() {
    f := BCalamocha卡拉莫查.(*卡拉莫查CalamochaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "calamocha",
		TitleName: "卡拉莫查",
		TitleCode: "b_calamocha",
	}
}
