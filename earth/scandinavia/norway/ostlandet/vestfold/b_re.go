package vestfold

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷ReBarony struct {
	feud.BaseBarony
}

var BRe雷 feud.Barony = &雷ReBarony{}

func init() {
    f := BRe雷.(*雷ReBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "re",
		TitleName: "雷",
		TitleCode: "b_re",
	}
}
