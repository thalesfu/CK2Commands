package mahdia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿加里卜AgarebBarony struct {
	feud.BaseBarony
}

var BAgareb阿加里卜 feud.Barony = &阿加里卜AgarebBarony{}

func init() {
    f := BAgareb阿加里卜.(*阿加里卜AgarebBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "agareb",
		TitleName: "阿加里卜",
		TitleCode: "b_agareb",
	}
}
