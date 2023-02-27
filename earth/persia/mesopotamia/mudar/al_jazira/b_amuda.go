package al_jazira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿穆达AmudaBarony struct {
	feud.BaseBarony
}

var BAmuda阿穆达 feud.Barony = &阿穆达AmudaBarony{}

func init() {
    f := BAmuda阿穆达.(*阿穆达AmudaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amuda",
		TitleName: "阿穆达",
		TitleCode: "b_amuda",
	}
}
