package kashgar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 措法TsofaBarony struct {
	feud.BaseBarony
}

var BTsofa措法 feud.Barony = &措法TsofaBarony{}

func init() {
    f := BTsofa措法.(*措法TsofaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsofa",
		TitleName: "措法",
		TitleCode: "b_tsofa",
	}
}
