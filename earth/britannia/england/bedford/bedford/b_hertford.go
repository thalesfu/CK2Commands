package bedford

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫特福德HertfordBarony struct {
	feud.BaseBarony
}

var BHertford赫特福德 feud.Barony = &赫特福德HertfordBarony{}

func init() {
    f := BHertford赫特福德.(*赫特福德HertfordBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hertford",
		TitleName: "赫特福德",
		TitleCode: "b_hertford",
	}
}
