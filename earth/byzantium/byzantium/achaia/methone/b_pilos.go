package methone

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮洛斯PilosBarony struct {
	feud.BaseBarony
}

var BPilos皮洛斯 feud.Barony = &皮洛斯PilosBarony{}

func init() {
    f := BPilos皮洛斯.(*皮洛斯PilosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pilos",
		TitleName: "皮洛斯",
		TitleCode: "b_pilos",
	}
}
