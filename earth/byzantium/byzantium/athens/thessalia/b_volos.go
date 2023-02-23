package thessalia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃洛斯VolosBarony struct {
	feud.BaseBarony
}

var BVolos沃洛斯 feud.Barony = &沃洛斯VolosBarony{}

func init() {
	f := BVolos沃洛斯.(*沃洛斯VolosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "volos",
		TitleName: "沃洛斯",
		TitleCode: "b_volos",
	}
}
