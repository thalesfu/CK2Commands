package al_bichri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代尔祖尔DeiralzurBarony struct {
	feud.BaseBarony
}

var BDeiralzur代尔祖尔 feud.Barony = &代尔祖尔DeiralzurBarony{}

func init() {
    f := BDeiralzur代尔祖尔.(*代尔祖尔DeiralzurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "deiralzur",
		TitleName: "代尔祖尔",
		TitleCode: "b_deiralzur",
	}
}
