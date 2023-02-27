package rashka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普里兹伦PrizrenBarony struct {
	feud.BaseBarony
}

var BPrizren普里兹伦 feud.Barony = &普里兹伦PrizrenBarony{}

func init() {
    f := BPrizren普里兹伦.(*普里兹伦PrizrenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "prizren",
		TitleName: "普里兹伦",
		TitleCode: "b_prizren",
	}
}
