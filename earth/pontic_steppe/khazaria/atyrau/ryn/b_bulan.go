package ryn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布兰BulanBarony struct {
	feud.BaseBarony
}

var BBulan布兰 feud.Barony = &布兰BulanBarony{}

func init() {
    f := BBulan布兰.(*布兰BulanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bulan",
		TitleName: "布兰",
		TitleCode: "b_bulan",
	}
}
