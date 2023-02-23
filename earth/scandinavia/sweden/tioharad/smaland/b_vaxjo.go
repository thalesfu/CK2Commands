package smaland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦克赫VaxjoBarony struct {
	feud.BaseBarony
}

var BVaxjo韦克赫 feud.Barony = &韦克赫VaxjoBarony{}

func init() {
	f := BVaxjo韦克赫.(*韦克赫VaxjoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vaxjo",
		TitleName: "韦克赫",
		TitleCode: "b_vaxjo",
	}
}
