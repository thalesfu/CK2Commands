package santiago

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维戈VigoBarony struct {
	feud.BaseBarony
}

var BVigo维戈 feud.Barony = &维戈VigoBarony{}

func init() {
    f := BVigo维戈.(*维戈VigoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vigo",
		TitleName: "维戈",
		TitleCode: "b_vigo",
	}
}
