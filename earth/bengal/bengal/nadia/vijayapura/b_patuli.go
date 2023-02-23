package vijayapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波堵梨PatuliBarony struct {
	feud.BaseBarony
}

var BPatuli波堵梨 feud.Barony = &波堵梨PatuliBarony{}

func init() {
	f := BPatuli波堵梨.(*波堵梨PatuliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "patuli",
		TitleName: "波堵梨",
		TitleCode: "b_patuli",
	}
}
