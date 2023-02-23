package nikaea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥普提马同OptimatumBarony struct {
	feud.BaseBarony
}

var BOptimatum奥普提马同 feud.Barony = &奥普提马同OptimatumBarony{}

func init() {
	f := BOptimatum奥普提马同.(*奥普提马同OptimatumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "optimatum",
		TitleName: "奥普提马同",
		TitleCode: "b_optimatum",
	}
}
