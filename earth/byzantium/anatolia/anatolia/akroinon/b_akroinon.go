package akroinon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克罗伊农AkroinonBarony struct {
	feud.BaseBarony
}

var BAkroinon阿克罗伊农 feud.Barony = &阿克罗伊农AkroinonBarony{}

func init() {
	f := BAkroinon阿克罗伊农.(*阿克罗伊农AkroinonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akroinon",
		TitleName: "阿克罗伊农",
		TitleCode: "b_akroinon",
	}
}
