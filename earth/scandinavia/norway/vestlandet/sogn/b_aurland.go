package sogn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾于兰AurlandBarony struct {
	feud.BaseBarony
}

var BAurland艾于兰 feud.Barony = &艾于兰AurlandBarony{}

func init() {
	f := BAurland艾于兰.(*艾于兰AurlandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aurland",
		TitleName: "艾于兰",
		TitleCode: "b_aurland",
	}
}
