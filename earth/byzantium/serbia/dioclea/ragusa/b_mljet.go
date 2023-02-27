package ragusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 姆列特MljetBarony struct {
	feud.BaseBarony
}

var BMljet姆列特 feud.Barony = &姆列特MljetBarony{}

func init() {
    f := BMljet姆列特.(*姆列特MljetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mljet",
		TitleName: "姆列特",
		TitleCode: "b_mljet",
	}
}
