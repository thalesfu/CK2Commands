package hendjan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔甘ArganBarony struct {
	feud.BaseBarony
}

var BArgan阿尔甘 feud.Barony = &阿尔甘ArganBarony{}

func init() {
	f := BArgan阿尔甘.(*阿尔甘ArganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "argan",
		TitleName: "阿尔甘",
		TitleCode: "b_argan",
	}
}
