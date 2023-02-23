package bumthang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 辛格SingyeBarony struct {
	feud.BaseBarony
}

var BSingye辛格 feud.Barony = &辛格SingyeBarony{}

func init() {
	f := BSingye辛格.(*辛格SingyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "singye",
		TitleName: "辛格",
		TitleCode: "b_singye",
	}
}
