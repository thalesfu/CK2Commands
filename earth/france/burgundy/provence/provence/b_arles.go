package provence

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔勒ArlesBarony struct {
	feud.BaseBarony
}

var BArles阿尔勒 feud.Barony = &阿尔勒ArlesBarony{}

func init() {
	f := BArles阿尔勒.(*阿尔勒ArlesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arles",
		TitleName: "阿尔勒",
		TitleCode: "b_arles",
	}
}
