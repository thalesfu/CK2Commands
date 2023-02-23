package abosyssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥胡斯AarhusBarony struct {
	feud.BaseBarony
}

var BAarhus奥胡斯 feud.Barony = &奥胡斯AarhusBarony{}

func init() {
	f := BAarhus奥胡斯.(*奥胡斯AarhusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aarhus",
		TitleName: "奥胡斯",
		TitleCode: "b_aarhus",
	}
}
