package tripuri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦罗诃GarhaBarony struct {
	feud.BaseBarony
}

var BGarha迦罗诃 feud.Barony = &迦罗诃GarhaBarony{}

func init() {
	f := BGarha迦罗诃.(*迦罗诃GarhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "garha",
		TitleName: "迦罗诃",
		TitleCode: "b_garha",
	}
}
