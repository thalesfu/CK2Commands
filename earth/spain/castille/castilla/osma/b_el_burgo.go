package osma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃尔布尔戈El_burgoBarony struct {
	feud.BaseBarony
}

var BEl_burgo埃尔布尔戈 feud.Barony = &埃尔布尔戈El_burgoBarony{}

func init() {
    f := BEl_burgo埃尔布尔戈.(*埃尔布尔戈El_burgoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "el_burgo",
		TitleName: "埃尔布尔戈",
		TitleCode: "b_el_burgo",
	}
}
