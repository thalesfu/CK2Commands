package wag

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 因德赫吉巴Indeher_gibaBarony struct {
	feud.BaseBarony
}

var BIndeher_giba因德赫吉巴 feud.Barony = &因德赫吉巴Indeher_gibaBarony{}

func init() {
    f := BIndeher_giba因德赫吉巴.(*因德赫吉巴Indeher_gibaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "indeher_giba",
		TitleName: "因德赫吉巴",
		TitleCode: "b_indeher_giba",
	}
}
