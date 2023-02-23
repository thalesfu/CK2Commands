package limisol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔西诺伊ArsinoeBarony struct {
	feud.BaseBarony
}

var BArsinoe阿尔西诺伊 feud.Barony = &阿尔西诺伊ArsinoeBarony{}

func init() {
	f := BArsinoe阿尔西诺伊.(*阿尔西诺伊ArsinoeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arsinoe",
		TitleName: "阿尔西诺伊",
		TitleCode: "b_arsinoe",
	}
}
