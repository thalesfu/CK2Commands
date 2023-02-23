package wight

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 本彻奇BonchurchBarony struct {
	feud.BaseBarony
}

var BBonchurch本彻奇 feud.Barony = &本彻奇BonchurchBarony{}

func init() {
	f := BBonchurch本彻奇.(*本彻奇BonchurchBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bonchurch",
		TitleName: "本彻奇",
		TitleCode: "b_bonchurch",
	}
}
