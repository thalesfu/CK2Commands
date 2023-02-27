package alampur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 汉比HampiBarony struct {
	feud.BaseBarony
}

var BHampi汉比 feud.Barony = &汉比HampiBarony{}

func init() {
    f := BHampi汉比.(*汉比HampiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hampi",
		TitleName: "汉比",
		TitleCode: "b_hampi",
	}
}
