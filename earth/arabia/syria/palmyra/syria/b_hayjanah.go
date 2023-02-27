package syria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海贾纳HayjanahBarony struct {
	feud.BaseBarony
}

var BHayjanah海贾纳 feud.Barony = &海贾纳HayjanahBarony{}

func init() {
    f := BHayjanah海贾纳.(*海贾纳HayjanahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hayjanah",
		TitleName: "海贾纳",
		TitleCode: "b_hayjanah",
	}
}
