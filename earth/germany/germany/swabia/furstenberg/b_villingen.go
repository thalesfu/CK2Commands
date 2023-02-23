package furstenberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲林根VillingenBarony struct {
	feud.BaseBarony
}

var BVillingen菲林根 feud.Barony = &菲林根VillingenBarony{}

func init() {
	f := BVillingen菲林根.(*菲林根VillingenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "villingen",
		TitleName: "菲林根",
		TitleCode: "b_villingen",
	}
}
