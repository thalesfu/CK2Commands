package zaozerye

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰维亚季内DevyatinyBarony struct {
	feud.BaseBarony
}

var BDevyatiny杰维亚季内 feud.Barony = &杰维亚季内DevyatinyBarony{}

func init() {
	f := BDevyatiny杰维亚季内.(*杰维亚季内DevyatinyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "devyatiny",
		TitleName: "杰维亚季内",
		TitleCode: "b_devyatiny",
	}
}
