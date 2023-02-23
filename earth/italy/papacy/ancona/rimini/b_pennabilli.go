package rimini

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 彭纳比利PennabilliBarony struct {
	feud.BaseBarony
}

var BPennabilli彭纳比利 feud.Barony = &彭纳比利PennabilliBarony{}

func init() {
	f := BPennabilli彭纳比利.(*彭纳比利PennabilliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pennabilli",
		TitleName: "彭纳比利",
		TitleCode: "b_pennabilli",
	}
}
