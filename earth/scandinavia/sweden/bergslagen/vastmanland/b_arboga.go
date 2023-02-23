package vastmanland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔博加ArbogaBarony struct {
	feud.BaseBarony
}

var BArboga阿尔博加 feud.Barony = &阿尔博加ArbogaBarony{}

func init() {
	f := BArboga阿尔博加.(*阿尔博加ArbogaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arboga",
		TitleName: "阿尔博加",
		TitleCode: "b_arboga",
	}
}
