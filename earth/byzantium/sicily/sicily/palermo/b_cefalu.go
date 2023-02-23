package palermo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切法卢CefaluBarony struct {
	feud.BaseBarony
}

var BCefalu切法卢 feud.Barony = &切法卢CefaluBarony{}

func init() {
	f := BCefalu切法卢.(*切法卢CefaluBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cefalu",
		TitleName: "切法卢",
		TitleCode: "b_cefalu",
	}
}
