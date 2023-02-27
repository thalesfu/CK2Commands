package wag

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨卡SaccaBarony struct {
	feud.BaseBarony
}

var BSacca萨卡 feud.Barony = &萨卡SaccaBarony{}

func init() {
    f := BSacca萨卡.(*萨卡SaccaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sacca",
		TitleName: "萨卡",
		TitleCode: "b_sacca",
	}
}
