package perche

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺让Nogent_le_rotrouBarony struct {
	feud.BaseBarony
}

var BNogent_le_rotrou诺让 feud.Barony = &诺让Nogent_le_rotrouBarony{}

func init() {
    f := BNogent_le_rotrou诺让.(*诺让Nogent_le_rotrouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nogent_le_rotrou",
		TitleName: "诺让",
		TitleCode: "b_nogent_le_rotrou",
	}
}
