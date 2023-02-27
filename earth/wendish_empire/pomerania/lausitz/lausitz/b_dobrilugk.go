package lausitz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多伯卢格DobrilugkBarony struct {
	feud.BaseBarony
}

var BDobrilugk多伯卢格 feud.Barony = &多伯卢格DobrilugkBarony{}

func init() {
    f := BDobrilugk多伯卢格.(*多伯卢格DobrilugkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dobrilugk",
		TitleName: "多伯卢格",
		TitleCode: "b_dobrilugk",
	}
}
