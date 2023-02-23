package puri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 补梨PuriBarony struct {
	feud.BaseBarony
}

var BPuri补梨 feud.Barony = &补梨PuriBarony{}

func init() {
	f := BPuri补梨.(*补梨PuriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "puri",
		TitleName: "补梨",
		TitleCode: "b_puri",
	}
}
