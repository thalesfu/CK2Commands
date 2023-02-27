package ankyra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳科莱亚NakoleiaBarony struct {
	feud.BaseBarony
}

var BNakoleia纳科莱亚 feud.Barony = &纳科莱亚NakoleiaBarony{}

func init() {
    f := BNakoleia纳科莱亚.(*纳科莱亚NakoleiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nakoleia",
		TitleName: "纳科莱亚",
		TitleCode: "b_nakoleia",
	}
}
