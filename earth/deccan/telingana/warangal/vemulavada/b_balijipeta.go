package vemulavada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆梨恃卑吒BalijipetaBarony struct {
	feud.BaseBarony
}

var BBalijipeta婆梨恃卑吒 feud.Barony = &婆梨恃卑吒BalijipetaBarony{}

func init() {
    f := BBalijipeta婆梨恃卑吒.(*婆梨恃卑吒BalijipetaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "balijipeta",
		TitleName: "婆梨恃卑吒",
		TitleCode: "b_balijipeta",
	}
}
