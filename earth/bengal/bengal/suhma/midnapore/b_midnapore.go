package midnapore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迷地尼城MidnaporeBarony struct {
	feud.BaseBarony
}

var BMidnapore迷地尼城 feud.Barony = &迷地尼城MidnaporeBarony{}

func init() {
	f := BMidnapore迷地尼城.(*迷地尼城MidnaporeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "midnapore",
		TitleName: "迷地尼城",
		TitleCode: "b_midnapore",
	}
}
