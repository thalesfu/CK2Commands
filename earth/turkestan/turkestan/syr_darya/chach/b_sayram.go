package chach

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞蓝SayramBarony struct {
	feud.BaseBarony
}

var BSayram塞蓝 feud.Barony = &塞蓝SayramBarony{}

func init() {
    f := BSayram塞蓝.(*塞蓝SayramBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sayram",
		TitleName: "塞蓝",
		TitleCode: "b_sayram",
	}
}
