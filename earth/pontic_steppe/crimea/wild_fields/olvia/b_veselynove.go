package olvia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦谢利诺韦VeselynoveBarony struct {
	feud.BaseBarony
}

var BVeselynove韦谢利诺韦 feud.Barony = &韦谢利诺韦VeselynoveBarony{}

func init() {
    f := BVeselynove韦谢利诺韦.(*韦谢利诺韦VeselynoveBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "veselynove",
		TitleName: "韦谢利诺韦",
		TitleCode: "b_veselynove",
	}
}
