package litomerice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰钦DecinBarony struct {
	feud.BaseBarony
}

var BDecin杰钦 feud.Barony = &杰钦DecinBarony{}

func init() {
    f := BDecin杰钦.(*杰钦DecinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "decin",
		TitleName: "杰钦",
		TitleCode: "b_decin",
	}
}
