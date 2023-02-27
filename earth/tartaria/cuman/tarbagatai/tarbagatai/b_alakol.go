package tarbagatai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉阔勒AlakolBarony struct {
	feud.BaseBarony
}

var BAlakol阿拉阔勒 feud.Barony = &阿拉阔勒AlakolBarony{}

func init() {
    f := BAlakol阿拉阔勒.(*阿拉阔勒AlakolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alakol",
		TitleName: "阿拉阔勒",
		TitleCode: "b_alakol",
	}
}
