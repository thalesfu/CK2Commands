package djerba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔古尔梅斯TaguermessBarony struct {
	feud.BaseBarony
}

var BTaguermess塔古尔梅斯 feud.Barony = &塔古尔梅斯TaguermessBarony{}

func init() {
	f := BTaguermess塔古尔梅斯.(*塔古尔梅斯TaguermessBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taguermess",
		TitleName: "塔古尔梅斯",
		TitleCode: "b_taguermess",
	}
}
