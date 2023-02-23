package songhay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昂松戈AnsongoBarony struct {
	feud.BaseBarony
}

var BAnsongo昂松戈 feud.Barony = &昂松戈AnsongoBarony{}

func init() {
	f := BAnsongo昂松戈.(*昂松戈AnsongoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ansongo",
		TitleName: "昂松戈",
		TitleCode: "b_ansongo",
	}
}
