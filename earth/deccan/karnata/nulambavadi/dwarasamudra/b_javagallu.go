package dwarasamudra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奢伐加卢JavagalluBarony struct {
	feud.BaseBarony
}

var BJavagallu奢伐加卢 feud.Barony = &奢伐加卢JavagalluBarony{}

func init() {
	f := BJavagallu奢伐加卢.(*奢伐加卢JavagalluBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "javagallu",
		TitleName: "奢伐加卢",
		TitleCode: "b_javagallu",
	}
}
