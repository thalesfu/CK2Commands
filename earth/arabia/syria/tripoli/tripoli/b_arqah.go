package tripoli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉赫ArqahBarony struct {
	feud.BaseBarony
}

var BArqah阿拉赫 feud.Barony = &阿拉赫ArqahBarony{}

func init() {
	f := BArqah阿拉赫.(*阿拉赫ArqahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arqah",
		TitleName: "阿拉赫",
		TitleCode: "b_arqah",
	}
}
