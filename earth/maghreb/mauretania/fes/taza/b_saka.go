package taza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨卡SakaBarony struct {
	feud.BaseBarony
}

var BSaka萨卡 feud.Barony = &萨卡SakaBarony{}

func init() {
	f := BSaka萨卡.(*萨卡SakaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saka",
		TitleName: "萨卡",
		TitleCode: "b_saka",
	}
}
