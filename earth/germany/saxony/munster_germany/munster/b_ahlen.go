package munster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿伦AhlenBarony struct {
	feud.BaseBarony
}

var BAhlen阿伦 feud.Barony = &阿伦AhlenBarony{}

func init() {
    f := BAhlen阿伦.(*阿伦AhlenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ahlen",
		TitleName: "阿伦",
		TitleCode: "b_ahlen",
	}
}
