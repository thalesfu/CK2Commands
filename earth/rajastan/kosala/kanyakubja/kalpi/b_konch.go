package kalpi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 军遮KonchBarony struct {
	feud.BaseBarony
}

var BKonch军遮 feud.Barony = &军遮KonchBarony{}

func init() {
	f := BKonch军遮.(*军遮KonchBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "konch",
		TitleName: "军遮",
		TitleCode: "b_konch",
	}
}
