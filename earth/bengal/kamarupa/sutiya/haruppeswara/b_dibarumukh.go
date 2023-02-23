package haruppeswara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提婆卢目佉DibarumukhBarony struct {
	feud.BaseBarony
}

var BDibarumukh提婆卢目佉 feud.Barony = &提婆卢目佉DibarumukhBarony{}

func init() {
	f := BDibarumukh提婆卢目佉.(*提婆卢目佉DibarumukhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dibarumukh",
		TitleName: "提婆卢目佉",
		TitleCode: "b_dibarumukh",
	}
}
