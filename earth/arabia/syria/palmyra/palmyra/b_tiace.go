package palmyra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提采TiaceBarony struct {
	feud.BaseBarony
}

var BTiace提采 feud.Barony = &提采TiaceBarony{}

func init() {
    f := BTiace提采.(*提采TiaceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tiace",
		TitleName: "提采",
		TitleCode: "b_tiace",
	}
}
