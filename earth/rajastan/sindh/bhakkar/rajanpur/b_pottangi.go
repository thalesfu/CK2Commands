package rajanpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菩檀耆PottangiBarony struct {
	feud.BaseBarony
}

var BPottangi菩檀耆 feud.Barony = &菩檀耆PottangiBarony{}

func init() {
	f := BPottangi菩檀耆.(*菩檀耆PottangiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pottangi",
		TitleName: "菩檀耆",
		TitleCode: "b_pottangi",
	}
}
