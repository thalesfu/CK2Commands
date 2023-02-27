package novgorod_seversky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴图林BatourynBarony struct {
	feud.BaseBarony
}

var BBatouryn巴图林 feud.Barony = &巴图林BatourynBarony{}

func init() {
    f := BBatouryn巴图林.(*巴图林BatourynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "batouryn",
		TitleName: "巴图林",
		TitleCode: "b_batouryn",
	}
}
