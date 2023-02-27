package balanjar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴兰贾尔BalanjarBarony struct {
	feud.BaseBarony
}

var BBalanjar巴兰贾尔 feud.Barony = &巴兰贾尔BalanjarBarony{}

func init() {
    f := BBalanjar巴兰贾尔.(*巴兰贾尔BalanjarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "balanjar",
		TitleName: "巴兰贾尔",
		TitleCode: "b_balanjar",
	}
}
