package ulm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰克TeckBarony struct {
	feud.BaseBarony
}

var BTeck泰克 feud.Barony = &泰克TeckBarony{}

func init() {
    f := BTeck泰克.(*泰克TeckBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "teck",
		TitleName: "泰克",
		TitleCode: "b_teck",
	}
}
