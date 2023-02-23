package calatrava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔马登AlmadeoBarony struct {
	feud.BaseBarony
}

var BAlmadeo阿尔马登 feud.Barony = &阿尔马登AlmadeoBarony{}

func init() {
	f := BAlmadeo阿尔马登.(*阿尔马登AlmadeoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "almadeo",
		TitleName: "阿尔马登",
		TitleCode: "b_almadeo",
	}
}
