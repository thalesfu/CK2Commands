package dorpat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基伦佩KirrumpaBarony struct {
	feud.BaseBarony
}

var BKirrumpa基伦佩 feud.Barony = &基伦佩KirrumpaBarony{}

func init() {
    f := BKirrumpa基伦佩.(*基伦佩KirrumpaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kirrumpa",
		TitleName: "基伦佩",
		TitleCode: "b_kirrumpa",
	}
}
