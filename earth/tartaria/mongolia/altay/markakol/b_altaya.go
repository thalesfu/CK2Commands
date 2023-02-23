package markakol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔塔亚AltayaBarony struct {
	feud.BaseBarony
}

var BAltaya阿尔塔亚 feud.Barony = &阿尔塔亚AltayaBarony{}

func init() {
	f := BAltaya阿尔塔亚.(*阿尔塔亚AltayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "altaya",
		TitleName: "阿尔塔亚",
		TitleCode: "b_altaya",
	}
}
