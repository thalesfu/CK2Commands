package taizz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 宰费尔ZafarBarony struct {
	feud.BaseBarony
}

var BZafar宰费尔 feud.Barony = &宰费尔ZafarBarony{}

func init() {
    f := BZafar宰费尔.(*宰费尔ZafarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zafar",
		TitleName: "宰费尔",
		TitleCode: "b_zafar",
	}
}
