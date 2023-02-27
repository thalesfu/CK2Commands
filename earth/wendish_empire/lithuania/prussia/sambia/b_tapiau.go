package sambia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔皮奥TapiauBarony struct {
	feud.BaseBarony
}

var BTapiau塔皮奥 feud.Barony = &塔皮奥TapiauBarony{}

func init() {
    f := BTapiau塔皮奥.(*塔皮奥TapiauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tapiau",
		TitleName: "塔皮奥",
		TitleCode: "b_tapiau",
	}
}
