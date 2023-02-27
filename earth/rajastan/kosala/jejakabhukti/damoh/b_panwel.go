package damoh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 般韦尔PanwelBarony struct {
	feud.BaseBarony
}

var BPanwel般韦尔 feud.Barony = &般韦尔PanwelBarony{}

func init() {
    f := BPanwel般韦尔.(*般韦尔PanwelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "panwel",
		TitleName: "般韦尔",
		TitleCode: "b_panwel",
	}
}
