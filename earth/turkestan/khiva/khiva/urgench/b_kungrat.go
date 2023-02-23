package urgench

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昆格勒KungratBarony struct {
	feud.BaseBarony
}

var BKungrat昆格勒 feud.Barony = &昆格勒KungratBarony{}

func init() {
	f := BKungrat昆格勒.(*昆格勒KungratBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kungrat",
		TitleName: "昆格勒",
		TitleCode: "b_kungrat",
	}
}
