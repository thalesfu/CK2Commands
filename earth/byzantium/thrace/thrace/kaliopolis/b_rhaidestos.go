package kaliopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱迪斯托斯RhaidestosBarony struct {
	feud.BaseBarony
}

var BRhaidestos莱迪斯托斯 feud.Barony = &莱迪斯托斯RhaidestosBarony{}

func init() {
    f := BRhaidestos莱迪斯托斯.(*莱迪斯托斯RhaidestosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rhaidestos",
		TitleName: "莱迪斯托斯",
		TitleCode: "b_rhaidestos",
	}
}
