package jarva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 努尔梅昆德NurmekundBarony struct {
	feud.BaseBarony
}

var BNurmekund努尔梅昆德 feud.Barony = &努尔梅昆德NurmekundBarony{}

func init() {
	f := BNurmekund努尔梅昆德.(*努尔梅昆德NurmekundBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nurmekund",
		TitleName: "努尔梅昆德",
		TitleCode: "b_nurmekund",
	}
}
