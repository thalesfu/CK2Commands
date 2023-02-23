package rashka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波洛格PologBarony struct {
	feud.BaseBarony
}

var BPolog波洛格 feud.Barony = &波洛格PologBarony{}

func init() {
	f := BPolog波洛格.(*波洛格PologBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "polog",
		TitleName: "波洛格",
		TitleCode: "b_polog",
	}
}
