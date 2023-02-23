package palencia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦凯伊VaccaeiBarony struct {
	feud.BaseBarony
}

var BVaccaei瓦凯伊 feud.Barony = &瓦凯伊VaccaeiBarony{}

func init() {
	f := BVaccaei瓦凯伊.(*瓦凯伊VaccaeiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vaccaei",
		TitleName: "瓦凯伊",
		TitleCode: "b_vaccaei",
	}
}
