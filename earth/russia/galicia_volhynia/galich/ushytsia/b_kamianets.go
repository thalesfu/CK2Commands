package ushytsia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡缅涅茨KamianetsBarony struct {
	feud.BaseBarony
}

var BKamianets卡缅涅茨 feud.Barony = &卡缅涅茨KamianetsBarony{}

func init() {
    f := BKamianets卡缅涅茨.(*卡缅涅茨KamianetsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kamianets",
		TitleName: "卡缅涅茨",
		TitleCode: "b_kamianets",
	}
}
