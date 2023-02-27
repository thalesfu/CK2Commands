package bereg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索伊沃SzolyvaBarony struct {
	feud.BaseBarony
}

var BSzolyva索伊沃 feud.Barony = &索伊沃SzolyvaBarony{}

func init() {
    f := BSzolyva索伊沃.(*索伊沃SzolyvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "szolyva",
		TitleName: "索伊沃",
		TitleCode: "b_szolyva",
	}
}
