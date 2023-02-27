package alexandria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆拉比特MaraboutBarony struct {
	feud.BaseBarony
}

var BMarabout穆拉比特 feud.Barony = &穆拉比特MaraboutBarony{}

func init() {
    f := BMarabout穆拉比特.(*穆拉比特MaraboutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marabout",
		TitleName: "穆拉比特",
		TitleCode: "b_marabout",
	}
}
