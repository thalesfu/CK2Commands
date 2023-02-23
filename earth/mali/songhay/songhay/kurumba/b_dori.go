package kurumba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多里DoriBarony struct {
	feud.BaseBarony
}

var BDori多里 feud.Barony = &多里DoriBarony{}

func init() {
	f := BDori多里.(*多里DoriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dori",
		TitleName: "多里",
		TitleCode: "b_dori",
	}
}
