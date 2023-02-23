package belgrade

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利波维茨LipovicBarony struct {
	feud.BaseBarony
}

var BLipovic利波维茨 feud.Barony = &利波维茨LipovicBarony{}

func init() {
	f := BLipovic利波维茨.(*利波维茨LipovicBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lipovic",
		TitleName: "利波维茨",
		TitleCode: "b_lipovic",
	}
}
