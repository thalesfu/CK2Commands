package taktse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 琼结QonggyaiBarony struct {
	feud.BaseBarony
}

var BQonggyai琼结 feud.Barony = &琼结QonggyaiBarony{}

func init() {
    f := BQonggyai琼结.(*琼结QonggyaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qonggyai",
		TitleName: "琼结",
		TitleCode: "b_qonggyai",
	}
}
