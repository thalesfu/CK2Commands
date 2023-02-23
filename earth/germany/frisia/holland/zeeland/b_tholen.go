package zeeland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托伦TholenBarony struct {
	feud.BaseBarony
}

var BTholen托伦 feud.Barony = &托伦TholenBarony{}

func init() {
	f := BTholen托伦.(*托伦TholenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tholen",
		TitleName: "托伦",
		TitleCode: "b_tholen",
	}
}
