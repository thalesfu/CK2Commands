package huelva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波蒂舒埃洛PortichueloBarony struct {
	feud.BaseBarony
}

var BPortichuelo波蒂舒埃洛 feud.Barony = &波蒂舒埃洛PortichueloBarony{}

func init() {
    f := BPortichuelo波蒂舒埃洛.(*波蒂舒埃洛PortichueloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "portichuelo",
		TitleName: "波蒂舒埃洛",
		TitleCode: "b_portichuelo",
	}
}
