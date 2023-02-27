package west_dvina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尔维BalviBarony struct {
	feud.BaseBarony
}

var BBalvi巴尔维 feud.Barony = &巴尔维BalviBarony{}

func init() {
    f := BBalvi巴尔维.(*巴尔维BalviBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "balvi",
		TitleName: "巴尔维",
		TitleCode: "b_balvi",
	}
}
