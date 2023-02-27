package bashkirs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别列别伊BelebeyBarony struct {
	feud.BaseBarony
}

var BBelebey别列别伊 feud.Barony = &别列别伊BelebeyBarony{}

func init() {
    f := BBelebey别列别伊.(*别列别伊BelebeyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belebey",
		TitleName: "别列别伊",
		TitleCode: "b_belebey",
	}
}
