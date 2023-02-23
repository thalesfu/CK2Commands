package djimi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥诺AunoBarony struct {
	feud.BaseBarony
}

var BAuno奥诺 feud.Barony = &奥诺AunoBarony{}

func init() {
	f := BAuno奥诺.(*奥诺AunoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "auno",
		TitleName: "奥诺",
		TitleCode: "b_auno",
	}
}
