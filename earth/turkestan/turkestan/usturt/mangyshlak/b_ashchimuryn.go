package mangyshlak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿希穆伦AshchimurynBarony struct {
	feud.BaseBarony
}

var BAshchimuryn阿希穆伦 feud.Barony = &阿希穆伦AshchimurynBarony{}

func init() {
	f := BAshchimuryn阿希穆伦.(*阿希穆伦AshchimurynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ashchimuryn",
		TitleName: "阿希穆伦",
		TitleCode: "b_ashchimuryn",
	}
}
