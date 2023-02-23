package kolomna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥热列利耶OzherelyeBarony struct {
	feud.BaseBarony
}

var BOzherelye奥热列利耶 feud.Barony = &奥热列利耶OzherelyeBarony{}

func init() {
	f := BOzherelye奥热列利耶.(*奥热列利耶OzherelyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ozherelye",
		TitleName: "奥热列利耶",
		TitleCode: "b_ozherelye",
	}
}
