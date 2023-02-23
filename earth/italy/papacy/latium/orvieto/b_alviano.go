package orvieto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔维亚诺AlvianoBarony struct {
	feud.BaseBarony
}

var BAlviano阿尔维亚诺 feud.Barony = &阿尔维亚诺AlvianoBarony{}

func init() {
	f := BAlviano阿尔维亚诺.(*阿尔维亚诺AlvianoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alviano",
		TitleName: "阿尔维亚诺",
		TitleCode: "b_alviano",
	}
}
