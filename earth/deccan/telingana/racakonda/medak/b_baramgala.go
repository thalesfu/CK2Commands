package medak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆兰迦罗BaramgalaBarony struct {
	feud.BaseBarony
}

var BBaramgala婆兰迦罗 feud.Barony = &婆兰迦罗BaramgalaBarony{}

func init() {
    f := BBaramgala婆兰迦罗.(*婆兰迦罗BaramgalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baramgala",
		TitleName: "婆兰迦罗",
		TitleCode: "b_baramgala",
	}
}
